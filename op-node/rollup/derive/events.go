package derive

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/hashicorp/go-multierror"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const EVENT_HOOK_ABI = `[
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
        }
      ],
      "name": "handle",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`



func parseLogs(receipts []*types.Receipt, eventHooks []bindings.EventHookRegistryEventHookItem) ([]*types.DepositTx, error) {
	var out []*types.DepositTx
	var result error
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for j, log := range rec.Logs {
			for _, hook := range eventHooks {
				if log.Address == hook.Origin && len(log.Topics) > 0 && log.Topics[0] == hook.Topic {
					fmt.Println("Found a matching topic")
					fmt.Println(log.Topics)
					fmt.Println(log.Data)

					registry, err := abi.JSON(strings.NewReader(EVENT_HOOK_ABI))
					if err != nil {
						fmt.Println("parseLogs json err", err)
						return  out, err
					}

					

						
					data, err := registry.Pack("handle", log.Topics[0], log.Address)
					if err != nil {
						fmt.Println("parseLogs pack err", err)
					}

					source := UserDepositSource{
							L1BlockHash: rec.BlockHash,
							LogIndex:    uint64(log.Index),
					}
					dep := types.DepositTx{
						From: L1InfoDepositerAddress,
						To: &hook.Receiver,
						Mint: big.NewInt(0),
						Value: big.NewInt(0),
						IsSystemTransaction: false,
						Gas: 300_000,
						SourceHash: source.SourceHash(),
						Data: data,
					}

					if err != nil {
						result = multierror.Append(result, fmt.Errorf("malformatted L1 deposit log in receipt %d, log %d: %w", i, j, err))
					} else {
						out = append(out, &dep)
					}
				}

			}
		}
	}
	return out, result
}

func DeriveEvents(receipts []*types.Receipt, eventHooks []bindings.EventHookRegistryEventHookItem) ([]hexutil.Bytes, error) {
	var result error
	userDeposits, err := parseLogs(receipts, eventHooks)
	if err != nil {
		result = multierror.Append(result, err)
	}
	encodedTxs := make([]hexutil.Bytes, 0, len(userDeposits))
	for i, tx := range userDeposits {
		opaqueTx, err := types.NewTx(tx).MarshalBinary()
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("failed to encode user tx %d", i))
		} else {
			encodedTxs = append(encodedTxs, opaqueTx)
		}
	}
	return encodedTxs, result
}
