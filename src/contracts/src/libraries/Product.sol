// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

library Models {
    struct Product {
        address seller;
        uint256 price;
        uint256 stock;
    }

    struct EscrowItem {
        uint256 itemId;
        address buyer;
        address seller;
        uint256 price;
        uint256 qty;
    }
}
