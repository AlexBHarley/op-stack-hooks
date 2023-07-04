package derive

import (
	// "bytes"
	// "encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-node/eth"
)

const (
    L1BurnFuncSignature = "report()"
    L1BurnArguments     = 0
    L1BurnLen           = 4 + 32*L1BurnArguments
)

var (
    L1BurnFuncBytes4 = crypto.Keccak256([]byte(L1BurnFuncSignature))[:4]
    L1BurnAddress    = common.HexToAddress("0x420000000000000000000000000000000000001b")
)


type L1BurnInfo struct {
    // Number uint64
    // Burn   uint64
}

func (info *L1BurnInfo) MarshalBinary() ([]byte, error) {
    data := make([]byte, L1BurnLen)
    offset := 0
    copy(data[offset:4], L1BurnFuncBytes4)
    // offset += 4
    // binary.BigEndian.PutUint64(data[offset+24:offset+32], info.Number)
    // offset += 32
    // binary.BigEndian.PutUint64(data[offset+24:offset+32], info.Burn)
    return data, nil
}

func (info *L1BurnInfo) UnmarshalBinary(data []byte) error {
    if len(data) != L1BurnLen {
        return fmt.Errorf("data is unexpected length: %d", len(data))
    }
    // var padding [24]byte
    // offset := 4
    // info.Number = binary.BigEndian.Uint64(data[offset+24 : offset+32])
    // if !bytes.Equal(data[offset:offset+24], padding[:]) {
    //     return fmt.Errorf("l1 burn tx number exceeds uint64 bounds: %x", data[offset:offset+32])
    // }
    // offset += 32
    // info.Burn = binary.BigEndian.Uint64(data[offset+24 : offset+32])
    // if !bytes.Equal(data[offset:offset+24], padding[:]) {
    //     return fmt.Errorf("l1 burn tx burn exceeds uint64 bounds: %x", data[offset:offset+32])
    // }
    return nil
}

func L1BurnDepositTxData(data []byte) (L1BurnInfo, error) {
    var info L1BurnInfo
    err := info.UnmarshalBinary(data)
    return info, err
}

func L1BurnDeposit(seqNumber uint64, block eth.BlockInfo, sysCfg eth.SystemConfig, regolith bool) (*types.DepositTx, error) {
    fmt.Println("L1BurnDeposit starting", block.NumberU64(), block.BaseFee().Uint64() * block.GasUsed())
    infoDat := L1BurnInfo{
        // Number: block.NumberU64(),
        // Burn:   block.BaseFee().Uint64() * block.GasUsed(),
    }
    data, err := infoDat.MarshalBinary()
    if err != nil {
        return nil, err
    }
    source := L1InfoDepositSource{
        L1BlockHash: block.Hash(),
        SeqNumber:   seqNumber,
    }

    fmt.Println("L1BurnDeposit done")
    out := &types.DepositTx{
        SourceHash:          source.SourceHash(),
        From:                L1InfoDepositerAddress,
        To:                  &L1BurnAddress,
        Mint:                nil,
        Value:               big.NewInt(0),
        Gas:                 150_000_000,
        IsSystemTransaction: true,
        Data:                data,
    }

    // With the regolith fork we disable the IsSystemTx functionality, and allocate real gas
	if regolith {
		out.IsSystemTransaction = false
		out.Gas = RegolithSystemTxGas
	}

    return out, nil
}

func L1BurnDepositBytes(seqNumber uint64, l1Info eth.BlockInfo, sysCfg eth.SystemConfig, regolith bool) ([]byte, error) {
    dep, err := L1BurnDeposit(seqNumber, l1Info, sysCfg, regolith)
    if err != nil {
        return nil, fmt.Errorf("failed to create L1 burn tx: %w", err)
    }
    l1Tx := types.NewTx(dep)
    opaqueL1Tx, err := l1Tx.MarshalBinary()
    if err != nil {
        return nil, fmt.Errorf("failed to encode L1 burn tx: %w", err)
    }
    return opaqueL1Tx, nil
}