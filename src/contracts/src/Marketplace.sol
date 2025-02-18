// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {Models} from "./libraries/Product.sol";
import {console} from "forge-std/console.sol";

contract Marketplace {
    error Marketplace_Unauthorized();
    error Marketplace_InvalidProduct();

    event ItemAdded(uint256 indexed id, address indexed seller);
    event PriceUpdated(uint256 indexed id, uint256 indexed price);
    event StockUpdated(uint256 indexed id, uint256 indexed newStock);

    uint256 private s_counter;
    mapping(uint256 => Models.Product) private s_products;
    address public s_escrowAddress;

    modifier onlyEscrowOrSeller(uint256 productId) {
        if (
            msg.sender != s_escrowAddress &&
            msg.sender != s_products[productId].seller
        ) {
            revert Marketplace_Unauthorized();
        }
        _;
    }

    function setEscrowAddress(address escrowAddress) external {
        s_escrowAddress = escrowAddress;
    }

    function addProduct(
        uint256 price,
        uint256 stock
    ) external returns (uint256) {
        s_counter++;
        uint256 productId = s_counter;
        s_products[productId] = Models.Product({
            seller: msg.sender,
            price: price,
            stock: stock
        });
        emit ItemAdded(productId, msg.sender);

        return productId;
    }

    function updateStock(
        uint256 id,
        uint256 stock
    ) external onlyEscrowOrSeller(id) {
        s_products[id].stock = stock;
        emit StockUpdated(id, stock);
    }

    function updatePrice(
        uint256 id,
        uint256 price
    ) external onlyEscrowOrSeller(id) {
        Models.Product memory product = s_products[id];
        if (msg.sender != s_products[id].seller) {
            revert Marketplace_Unauthorized();
        }
        s_products[id].price = price;

        emit PriceUpdated(id, price);
    }

    function getProduct(
        uint256 id
    ) external view returns (address, uint256, uint256) {
        Models.Product memory product = s_products[id];
        if (product.seller == address(0)) {
            revert Marketplace_InvalidProduct();
        }
        return (product.seller, product.price, product.stock);
    }
}
