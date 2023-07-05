// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IEventHook {
    function handle(
        address origin,
        bytes calldata topics,
        bytes calldata data
    ) external;
}
