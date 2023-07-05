// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import { Semver } from "../../universal/Semver.sol";

struct EventHookItem {
    bytes32 topic;
    address origin;
    address receiver;
}

/// @custom:proxied
/// @custom:predeploy 0x420000000000000000000000000000000000001c
/// @title EventHookRegistry
/// @notice A registry of hooks, that handle events from L1. A hook is defined as a
/// topic and address combination that the rollup should listen for logs from, as well
/// as a destination contract that should be invoked with the event data.
contract EventHookRegistry is Ownable {
    using EnumerableSet for EnumerableSet.Bytes32Set;

    mapping(bytes32 => EventHookItem) private hooks;
    EnumerableSet.Bytes32Set private hookIds;

    event EventHookAdded(bytes32 indexed id, bytes32 topic, address origin, address receiver);
    event EventHookRemoved(bytes32 indexed id);

    constructor(address _owner) Ownable() {
        _transferOwnership(_owner);
    }

    function addEventHook(
        bytes32 _topic,
        address _origin,
        address _receiver
    ) external onlyOwner returns (bytes32) {
        bytes32 id = keccak256(abi.encode(_topic, _origin, _receiver));
        require(!hookIds.contains(id), "hook exists");

        EventHookItem memory hook = EventHookItem({
            topic: _topic,
            origin: _origin,
            receiver: _receiver
        });
        hooks[id] = hook;
        hookIds.add(id);

        emit EventHookAdded(id, _topic, _origin, _receiver);

        return id;
    }

    function removeEventHook(bytes32 id) external onlyOwner {
        require(hookIds.contains(id), "!hook");

        delete hooks[id];
        hookIds.remove(id);

        emit EventHookRemoved(id);
    }

    function getEventHooks() external view returns (EventHookItem[] memory) {
        EventHookItem[] memory items = new EventHookItem[](hookIds.length());

        for (uint256 i = 0; i < hookIds.length(); i++) {
            bytes32 id = hookIds.at(i);
            items[i] = hooks[id];
        }

        return items;
    }
}
