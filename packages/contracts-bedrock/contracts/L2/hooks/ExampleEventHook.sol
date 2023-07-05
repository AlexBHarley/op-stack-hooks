// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { AbstractEventHook } from "./AbstractEventHook.sol";

/// @title ExampleEventHook
/// @notice Pair this with the ExampleEventEmitter contract
/// @notice Receives an event from the ExampleEventEmitter and .
contract ExampleEventHook is AbstractEventHook {
    bytes32 internal lastHandledTopic;
    address internal lastHandledAddress;

    constructor() {}

    function handle(
        address origin,
        bytes calldata topics,
        bytes calldata data
    ) external onlyDepositor {
        (address origin, address sender, uint256 count) = abi.decode(
            topics,
            (address, address, uint256)
        );
    }

    function getLastHandledTopic() external view returns (bytes32) {
        return lastHandledTopic;
    }

    function getLastHandledAddress() external view returns (address) {
        return lastHandledAddress;
    }
}
