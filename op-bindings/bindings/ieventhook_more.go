// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const IEventHookStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var IEventHookStorageLayout = new(solc.StorageLayout)

var IEventHookDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(IEventHookStorageLayoutJSON), IEventHookStorageLayout); err != nil {
		panic(err)
	}

	layouts["IEventHook"] = IEventHookStorageLayout
	deployedBytecodes["IEventHook"] = IEventHookDeployedBin
}
