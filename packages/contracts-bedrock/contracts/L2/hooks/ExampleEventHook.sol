// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { AbstractEventHook } from "./AbstractEventHook.sol";

/// @title ExampleEventHook
/// @notice Pair this with the ExampleEventEmitter contract
/// @notice Receives an event from the ExampleEventEmitter and writes some
/// values to storage.
contract ExampleEventHook is AbstractEventHook {
    address internal lastOriginHandled;

    address internal lastParsedOrigin;
    address internal lastParsedSender;
    uint256 internal lastParsedCount;

    constructor() {}

    function handle(
        address origin,
        bytes calldata topics,
        bytes calldata data
    ) external onlyDepositor {
        (bytes32 topic0, bytes32 topic1, bytes32 _topic2, bytes32 _topic3) = abi.decode(
            topics,
            (bytes32, bytes32, bytes32, bytes32)
        );

        lastOriginHandled = origin;

        lastParsedOrigin = address(bytes20(topic0));
        lastParsedSender = address(bytes20(topic1));
        lastParsedCount = uint256(bytes32(data));
    }

    function getLastOriginHandled() external view returns (address) {
        return lastOriginHandled;
    }

    function getLastParsedOrigin() external view returns (address) {
        return lastParsedOrigin;
    }

    function getLastParsedSender() external view returns (address) {
        return lastParsedSender;
    }

    function getLastParsedCount() external view returns (uint256) {
        return lastParsedCount;
    }
}
