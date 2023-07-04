// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { EnumerableMap } from "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

interface EventHook {
    function handle(bytes32, address) external;
}

/// @custom:proxied
/// @custom:predeploy 0x420000000000000000000000000000000000001c
/// @title EventHookRegistry
/// @notice blah blah.
contract EventHookRegistry {
    using EnumerableMap for EnumerableMap.UintToAddressMap;

    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;
    address owner;

    string public constant hello = "hello";

    struct EventHookItem {
        bytes32 topic;
        address origin;
        address receiver;
    }

    EventHookItem[] public hooks;

    constructor() {
        // owner = _owner;
    }

    function goodbye() external returns (string memory) {
        return "goodbye";
    }

    function addEventHook(
        bytes32 _topic,
        address _origin,
        address _receiver
    ) external {
        hooks.push(EventHookItem({ topic: _topic, origin: _origin, receiver: _receiver }));
    }

    function getEventHooks() external view returns (EventHookItem[] memory) {
        EventHookItem[] memory items = new EventHookItem[](hooks.length);

        for (uint256 i = 0; i < hooks.length; i++) {
            items[i] = hooks[i];
        }

        return items;
    }
}
