// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const BurnStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var BurnStorageLayout = new(solc.StorageLayout)

var BurnDeployedBin = "0x73000000000000000000000000000000000000000030146080604052600080fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(BurnStorageLayoutJSON), BurnStorageLayout); err != nil {
		panic(err)
	}

	layouts["Burn"] = BurnStorageLayout
	deployedBytecodes["Burn"] = BurnDeployedBin
}
