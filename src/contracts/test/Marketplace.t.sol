// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Test} from "forge-std/Test.sol";
import {Marketplace} from "../src/Marketplace.sol";
import {Escrow} from "../src/Escrow.sol";
import {DeployContracts} from "../script/DeployContracts.s.sol";
import {console} from "forge-std/console.sol";

contract MarketplaceTest is Test {
    Marketplace marketplace;
    Escrow escrow;

    function setUp() public {
        DeployContracts deployer = new DeployContracts();
        (marketplace, escrow) = deployer.run();
    }

    function testShouldSetEscrowAddress() public {
        assertEq(address(escrow), marketplace.s_escrowAddress());
    }

    function testAddProductShouldReturnCreatedId() public {
        uint256 productId = marketplace.addProduct(100, 10);
        assertEq(productId, 1);
    }

    function testAddProductShouldAddTheProduct() public {
        address user = address(123);
        vm.prank(user);
        vm.deal(user, 1 ether);
        uint256 productId = marketplace.addProduct(100, 10);

        (address seller, uint256 price, uint256 stock) = marketplace.getProduct(
            productId
        );

        assertEq(seller, address(123));
        assertEq(price, 100);
        assertEq(stock, 10);
    }

    function testUpdateStockShouldReturnErrorForUnauthorized() public {
        address user = address(2345);
        vm.prank(user);
        marketplace.addProduct(100, 10);
        vm.expectRevert();
        marketplace.updateStock(1, 5);
    }

    function testUpdateStockShouldUpdateStockForEscrow() public {
        marketplace.addProduct(100, 10);
        vm.prank(address(escrow));
        marketplace.updateStock(1, 5);

        (, , uint256 stock) = marketplace.getProduct(1);
        assertEq(stock, 5);
    }

    function testUpdateStockShouldUpdateStockForSeller() public {
        marketplace.addProduct(100, 10);
        marketplace.updateStock(1, 5);
        (, , uint256 stock) = marketplace.getProduct(1);
        assertEq(stock, 5);
    }

    function testUpdatePriceShouldReturnErrorForUnauthorized() public {
        address user = address(234);
        vm.prank(user);
        marketplace.addProduct(100, 10);
        vm.expectRevert();
        marketplace.updatePrice(1, 0.5 ether);
    }

    function testUpdatePriceShouldUpdatePriceForEscrow() public {
        vm.prank(address(escrow));
        marketplace.addProduct(100, 10);

        vm.prank(address(escrow));
        marketplace.updatePrice(1, 0.5 ether);

        (, uint256 price, ) = marketplace.getProduct(1);
        assertEq(price, 0.5 ether);
    }

    function testUpdatePriceShouldUpdatePriceForSeller() public {
        marketplace.addProduct(100, 10);
        marketplace.updatePrice(1, 0.5 ether);
        (, uint256 price, ) = marketplace.getProduct(1);
        assertEq(price, 0.5 ether);
    }

    function testGetProductShouldReturnErrorForNonExistentProduct() public {
        vm.expectRevert();
        marketplace.getProduct(1);
    }

    function testGetProductShouldReturnProduct() public {
        marketplace.addProduct(100, 10);
        (address seller, uint256 price, uint256 stock) = marketplace.getProduct(
            1
        );
        assertEq(seller, address(this));
        assertEq(price, 100);
        assertEq(stock, 10);
    }
}
