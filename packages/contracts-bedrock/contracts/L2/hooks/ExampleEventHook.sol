import { IEventHook } from "./IEventHook.sol";

contract ExampleEventHook is IEventHook {
    bytes32 internal lastHandledTopic;
    address internal lastHandledAddress;

    constructor() {}

    function handle(bytes32 _topic, address _address) external {
        lastHandledTopic = _topic;
        lastHandledAddress = _address;
    }

    function getLastHandledTopic() external view returns (bytes32) {
        return lastHandledTopic;
    }

    function getLastHandledAddress() external view returns (address) {
        return lastHandledAddress;
    }
}
