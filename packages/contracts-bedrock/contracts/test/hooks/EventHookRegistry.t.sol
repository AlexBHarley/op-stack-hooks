// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Test } from "forge-std/Test.sol";

import { Proxy } from "../../universal/Proxy.sol";
import { EventHookRegistry, EventHookItem } from "../../L2/hooks/EventHookRegistry.sol";

contract EventHookRegistryTest is Test {
    EventHookRegistry internal registry;

    bytes32 internal topic = bytes32("");
    address internal origin = address(1);
    address internal receiver = address(1);
    bytes32 internal id = keccak256(abi.encode(topic, origin, receiver));

    event EventHookAdded(bytes32 indexed id, bytes32 topic, address origin, address receiver);
    event EventHookRemoved(bytes32 indexed id);

    function setUp() public virtual {
        registry = new EventHookRegistry(address(this));
    }

    function testRegisteringEvent() public {
        vm.expectEmit(true, true, true, true);
        emit EventHookAdded(id, topic, origin, receiver);

        registry.addEventHook(topic, origin, receiver);

        EventHookItem[] memory hooks = registry.getEventHooks();
        assertEq(hooks.length, 1);
        assertEq(hooks[0].topic, topic);
        assertEq(hooks[0].origin, origin);
        assertEq(hooks[0].receiver, receiver);
    }

    function testRemovingEvent() public {
        registry.addEventHook(topic, origin, receiver);

        vm.expectEmit(true, true, true, true);
        emit EventHookRemoved(id);

        registry.removeEventHook(id);

        EventHookItem[] memory hooks = registry.getEventHooks();
        assertEq(hooks.length, 0);
    }

    function testOnlyOwner() public {
        vm.prank(address(1234));
        vm.expectRevert(bytes("Ownable: caller is not the owner"));
        registry.addEventHook(topic, origin, receiver);

        vm.prank(address(1234));
        vm.expectRevert(bytes("Ownable: caller is not the owner"));
        registry.removeEventHook(id);
    }
}
