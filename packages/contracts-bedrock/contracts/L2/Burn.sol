// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @custom:proxied
/// @custom:predeploy 0x420000000000000000000000000000000000001b
/// @title Burn
/// @notice blah blah.
contract Burn {
    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

    /**
     * @notice Total amount of ETH burned on L1.
     */
    uint256 public total = 1;

    /**
     * @notice Mapping of blocks numbers to total burn.
     */
    mapping(uint64 => uint256) public reports;

    /**
     * @notice Allows the system address to submit a report.
     *
     */
    function report() external {
        total += 1;
        // require(
        //     msg.sender == 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001,
        //     "L1Burn: reports can only be made from system address"
        // );

        // total += _burn;
        // reports[_blocknum] = total;
    }

    // /**
    //  * @notice Tallies up the total burn since a given block number.
    //  *
    //  * @param _blocknum L1 block number to tally from.
    //  *
    //  * @return Total amount of ETH burned since the given block number;
    //  */
    // function tally(uint64 _blocknum) external view returns (uint256) {
    //     return total - reports[_blocknum];
    // }
}
