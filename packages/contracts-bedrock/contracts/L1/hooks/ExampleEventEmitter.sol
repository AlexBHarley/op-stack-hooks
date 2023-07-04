contract ExampleEventEmitter {
    // Topic: 0x12210f92675543a3eee7d9f6cc64eaca8eb1431502f685da3f48e7593e2b7f1e
    event ExampleEvent(uint256);

    uint256 internal counter;

    constructor() {}

    function emitEvent() external {
        counter++;

        emit ExampleEvent(counter);
    }
}
