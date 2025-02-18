// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script} from "forge-std/Script.sol";
import {Marketplace} from "../src/Marketplace.sol";
import {Escrow} from "../src/Escrow.sol";

contract DeployContracts is Script {
    function run() public returns (Marketplace, Escrow) {
        vm.startBroadcast();

        Marketplace marketplace = new Marketplace();

        Escrow escrow = new Escrow(address(marketplace));

        marketplace.setEscrowAddress(address(escrow));

        vm.stopBroadcast();
        return (marketplace, escrow);
    }
}
