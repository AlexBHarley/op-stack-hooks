// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IEventHook } from "./IEventHook.sol";

abstract contract AbstractEventHook is IEventHook {
    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

    modifier onlyDepositor() {
        require(msg.sender == DEPOSITOR_ACCOUNT, "!depositor");
        _;
    }
}
