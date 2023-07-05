package sources

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/sources/caching"
)

type L2ClientConfig struct {
	EthClientConfig

	L2BlockRefsCacheSize int
	L1ConfigsCacheSize   int

	RollupCfg *rollup.Config
}

func L2ClientDefaultConfig(config *rollup.Config, trustRPC bool) *L2ClientConfig {
	// Cache 3/2 worth of sequencing window of payloads, block references, receipts and txs
	span := int(config.SeqWindowSize) * 3 / 2
	// Estimate number of L2 blocks in this span of L1 blocks
	// (there's always one L2 block per L1 block, L1 is thus the minimum, even if block time is very high)
	if config.BlockTime < 12 && config.BlockTime > 0 {
		span *= 12
		span /= int(config.BlockTime)
	}
	fullSpan := span
	if span > 1000 { // sanity cap. If a large sequencing window is configured, do not make the cache too large
		span = 1000
	}
	return &L2ClientConfig{
		EthClientConfig: EthClientConfig{
			// receipts and transactions are cached per block
			ReceiptsCacheSize:     span,
			TransactionsCacheSize: span,
			HeadersCacheSize:      span,
			PayloadsCacheSize:     span,
			MaxRequestsPerBatch:   20, // TODO: tune batch param
			MaxConcurrentRequests: 10,
			TrustRPC:              trustRPC,
			MustBePostMerge:       true,
			RPCProviderKind:       RPCKindBasic,
			MethodResetDuration:   time.Minute,
		},
		// Not bounded by span, to cover find-sync-start range fully for speedy recovery after errors.
		L2BlockRefsCacheSize: fullSpan,
		L1ConfigsCacheSize:   span,
		RollupCfg:            config,
	}
}

// L2Client extends EthClient with functions to fetch and cache eth.L2BlockRef values.
type L2Client struct {
	*EthClient
	rollupCfg *rollup.Config

	// cache L2BlockRef by hash
	// common.Hash -> eth.L2BlockRef
	l2BlockRefsCache *caching.LRUCache

	// cache SystemConfig by L2 hash
	// common.Hash -> eth.SystemConfig
	systemConfigsCache *caching.LRUCache
}

// NewL2Client constructs a new L2Client instance. The L2Client is a thin wrapper around the EthClient with added functions
// for fetching and caching eth.L2BlockRef values. This includes fetching an L2BlockRef by block number, label, or hash.
// See: [L2BlockRefByLabel], [L2BlockRefByNumber], [L2BlockRefByHash]
func NewL2Client(client client.RPC, log log.Logger, metrics caching.Metrics, config *L2ClientConfig) (*L2Client, error) {
	ethClient, err := NewEthClient(client, log, metrics, &config.EthClientConfig)
	if err != nil {
		return nil, err
	}

	return &L2Client{
		EthClient:          ethClient,
		rollupCfg:          config.RollupCfg,
		l2BlockRefsCache:   caching.NewLRUCache(metrics, "blockrefs", config.L2BlockRefsCacheSize),
		systemConfigsCache: caching.NewLRUCache(metrics, "systemconfigs", config.L1ConfigsCacheSize),
	}, nil
}

// L2BlockRefByLabel returns the [eth.L2BlockRef] for the given block label.
func (s *L2Client) L2BlockRefByLabel(ctx context.Context, label eth.BlockLabel) (eth.L2BlockRef, error) {
	payload, err := s.PayloadByLabel(ctx, label)
	if err != nil {
		// Both geth and erigon like to serve non-standard errors for the safe and finalized heads, correct that.
		// This happens when the chain just started and nothing is marked as safe/finalized yet.
		if strings.Contains(err.Error(), "block not found") || strings.Contains(err.Error(), "Unknown block") {
			err = ethereum.NotFound
		}
		// w%: wrap to preserve ethereum.NotFound case
		return eth.L2BlockRef{}, fmt.Errorf("failed to determine L2BlockRef of %s, could not get payload: %w", label, err)
	}
	ref, err := derive.PayloadToBlockRef(payload, &s.rollupCfg.Genesis)
	if err != nil {
		return eth.L2BlockRef{}, err
	}
	s.l2BlockRefsCache.Add(ref.Hash, ref)
	return ref, nil
}

// L2BlockRefByNumber returns the [eth.L2BlockRef] for the given block number.
func (s *L2Client) L2BlockRefByNumber(ctx context.Context, num uint64) (eth.L2BlockRef, error) {
	payload, err := s.PayloadByNumber(ctx, num)
	if err != nil {
		// w%: wrap to preserve ethereum.NotFound case
		return eth.L2BlockRef{}, fmt.Errorf("failed to determine L2BlockRef of height %v, could not get payload: %w", num, err)
	}
	ref, err := derive.PayloadToBlockRef(payload, &s.rollupCfg.Genesis)
	if err != nil {
		return eth.L2BlockRef{}, err
	}
	s.l2BlockRefsCache.Add(ref.Hash, ref)
	return ref, nil
}

// L2BlockRefByHash returns the [eth.L2BlockRef] for the given block hash.
// The returned BlockRef may not be in the canonical chain.
func (s *L2Client) L2BlockRefByHash(ctx context.Context, hash common.Hash) (eth.L2BlockRef, error) {
	if ref, ok := s.l2BlockRefsCache.Get(hash); ok {
		return ref.(eth.L2BlockRef), nil
	}


	payload, err := s.PayloadByHash(ctx, hash)
	if err != nil {
		// w%: wrap to preserve ethereum.NotFound case
		return eth.L2BlockRef{}, fmt.Errorf("failed to determine block-hash of hash %v, could not get payload: %w", hash, err)
	}
	ref, err := derive.PayloadToBlockRef(payload, &s.rollupCfg.Genesis)
	if err != nil {
		return eth.L2BlockRef{}, err
	}
	s.l2BlockRefsCache.Add(ref.Hash, ref)
	return ref, nil
}

// SystemConfigByL2Hash returns the [eth.SystemConfig] (matching the config updates up to and including the L1 origin) for the given L2 block hash.
// The returned [eth.SystemConfig] may not be in the canonical chain when the hash is not canonical.
func (s *L2Client) SystemConfigByL2Hash(ctx context.Context, hash common.Hash) (eth.SystemConfig, error) {
	if ref, ok := s.systemConfigsCache.Get(hash); ok {
		return ref.(eth.SystemConfig), nil
	}

	payload, err := s.PayloadByHash(ctx, hash)
	if err != nil {
		// w%: wrap to preserve ethereum.NotFound case
		return eth.SystemConfig{}, fmt.Errorf("failed to determine block-hash of hash %v, could not get payload: %w", hash, err)
	}
	cfg, err := derive.PayloadToSystemConfig(payload, s.rollupCfg)
	if err != nil {
		return eth.SystemConfig{}, err
	}
	s.systemConfigsCache.Add(hash, cfg)
	return cfg, nil
}

func (s *L2Client) L2EventHooks(ctx context.Context, blockNumber uint64) ([]bindings.EventHookItem, error) {
	fmt.Println("L2EventHooks start")
	hookRegistryAbi := `[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [],
      "name": "DEPOSITOR_ACCOUNT",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "_topic",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "_origin",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_receiver",
          "type": "address"
        }
      ],
      "name": "addEventHook",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getEventHooks",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bytes32",
              "name": "topic",
              "type": "bytes32"
            },
            {
              "internalType": "address",
              "name": "origin",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "receiver",
              "type": "address"
            }
          ],
          "internalType": "struct EventHookRegistry.EventHookItem[]",
          "name": "",
          "type": "tuple[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "goodbye",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "hello",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "hooks",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "topic",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "origin",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "receiver",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]`

	registry, err := abi.JSON(strings.NewReader(hookRegistryAbi))
	if err != nil {
		fmt.Println("json err", err)
		return nil, err
	}

	data, err := registry.Pack("getEventHooks")
	if err != nil {
		fmt.Println("pack err", err)
		return nil, err
	}

	fmt.Println("data", data)
	fmt.Println("blockNumber", blockNumber)

	to := common.HexToAddress("0x420000000000000000000000000000000000001c")
	call := map[string]interface{}{
		"to": &to,
		"data": hexutil.Encode(data),
		"from": to,
	}

	fmt.Println("call", call)

	var hex hexutil.Bytes
	err = s.client.CallContext(ctx, &hex, "eth_call", call, hexutil.EncodeUint64(blockNumber));
	if err != nil {
		fmt.Println("CallContext", err)
		return nil, err
	}

	var out []interface{}
	err = registry.UnpackIntoInterface(&out, "getEventHooks", hex)
	if err != nil {
		fmt.Println("UnpackIntoInterface err", err)
		return  nil, err
	}

	outActual := make([]bindings.EventHookItem, len(out))
	for i, item := range out {
		outActual[i] = *abi.ConvertType(item, new(bindings.EventHookItem)).(*bindings.EventHookItem)
	}
		
	fmt.Println("Got", len(outActual), "event hooks")
	
	return outActual, nil
}
