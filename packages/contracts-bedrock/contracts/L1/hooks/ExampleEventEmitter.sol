// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title ExampleEventEmitter
/// @notice Pair this with the ExampleEventHook contract
/// @notice Emits an event with an ever increasing counter.
contract ExampleEventEmitter {
    uint256 internal counter;

    // 0x906c0fc836aaccaf76ef8a4168843fbbb1d6fb940e6e045ed4d32c1bcefbc7c8
    event ExampleEvent(address indexed _contract, address indexed sender, uint256 count);

    constructor() {}

    function emitEvent() external {
        counter++;

        emit ExampleEvent(address(this), msg.sender, counter);
    }
}
