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

const EVENT_HOOK_ABI = `
[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "origin",
				"type": "address"
			},
			{
				"internalType": "bytes",
				"name": "topics",
				"type": "bytes"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "handle",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

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

	registry, err := abi.JSON(strings.NewReader(EVENT_HOOK_ABI))
	if err != nil {
		fmt.Println("buildHookCallData json err", err)
		return nil, nil
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

	result, err := registry.Pack("handle", log.Address, packedTopics, log.Data)
	if err != nil {
		fmt.Println("buildHookCallData pack err", err)
		return nil, nil
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
					fmt.Println("Found a matching topic", log.Topics[0])

					callData, err := buildHookCallData(log)

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
