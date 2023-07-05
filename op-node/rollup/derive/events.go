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

// Here we convert a log into a handle(origin, topics, data) call. We
// pack the logs (and normalise to 4 logs) so they can be safely abi.Decoded
// in Solidity land
func buildHookCallData(log *types.Log) ([]byte, error) {
	var logTopics [4]interface{}
	for i := 0; i < 4; i++ {
		var by [32]byte
		if i < len(log.Topics) {
			val := log.Topics[i]
			copy(by[:], val.Bytes())
		} else {
			by = [32]byte{}
		}
		logTopics[i] = by
	}

	registry, err := abi.JSON(strings.NewReader(bindings.IEventHookABI))
	if err != nil {
		return nil, err
	}

	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
	arguments := abi.Arguments{
		{Type: bytes32Ty},
		{Type: bytes32Ty},
		{Type: bytes32Ty},
		{Type: bytes32Ty},
	}

	packedTopics, err := arguments.Pack(
		logTopics[0],
		logTopics[1],
		logTopics[2],
		logTopics[3],
	)
	if err != nil {
		return nil, err
	}

	result, err := registry.Pack("handle", log.Address, packedTopics, log.Data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
 

func parseLogs(receipts []*types.Receipt, eventHooks []bindings.EventHookItem) ([]*types.DepositTx, error) {
	var out []*types.DepositTx
	var result error
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for j, log := range rec.Logs {
			for _, hook := range eventHooks {
				if log.Address == hook.Origin && len(log.Topics) > 0 && log.Topics[0] == hook.Topic {
					callData, err := buildHookCallData(log)
					if err != nil {
						result = multierror.Append(result, fmt.Errorf("unable to build hook calldata %d, log %d: %w", i, j, err))
						continue
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
						Gas: 500_000,
						SourceHash: source.SourceHash(),
						Data: callData,
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

func DeriveEvents(receipts []*types.Receipt, eventHooks []bindings.EventHookItem) ([]hexutil.Bytes, error) {
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
