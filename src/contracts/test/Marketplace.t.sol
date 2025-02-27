// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Test} from "forge-std/Test.sol";
import {Marketplace} from "../src/Marketplace.sol";
import {Escrow} from "../src/Escrow.sol";
import {DeployContracts} from "../script/DeployContracts.s.sol";

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
        uint256 productId = marketplace.addProduct(100, 10, "c206a740-3904-4b79-8c9c-36a9d666ff76");
        assertEq(productId, 1);
    }

    function testAddProductShouldAddTheProduct() public {
        address user = address(123);
        vm.prank(user);
        vm.deal(user, 1 ether);
        uint256 productId = marketplace.addProduct(100, 10, "070c2d10-5562-43a4-98c3-0787518c07df");

        (address seller, uint256 price, uint256 stock) = marketplace.getProduct(productId);

        assertEq(seller, address(123));
        assertEq(price, 100);
        assertEq(stock, 10);
    }

    function testUpdateStockShouldReturnErrorForUnauthorized() public {
        address user = address(2345);
        vm.prank(user);
        marketplace.addProduct(100, 10, "54d738a7-f3f3-4b75-bc78-01c5c539fed4");
        vm.expectRevert();
        marketplace.updateStock(1, 5);
    }

    function testUpdateStockShouldUpdateStockForEscrow() public {
        marketplace.addProduct(100, 10, "88aaf7bf-4b76-4505-9d62-9a251ec25133");
        vm.prank(address(escrow));
        marketplace.updateStock(1, 5);

        (,, uint256 stock) = marketplace.getProduct(1);
        assertEq(stock, 5);
    }

    function testUpdateStockShouldUpdateStockForSeller() public {
        marketplace.addProduct(100, 10, "76aa2f2c-6a2c-4360-99ac-d29d4cc5b05a");
        marketplace.updateStock(1, 5);
        (,, uint256 stock) = marketplace.getProduct(1);
        assertEq(stock, 5);
    }

    function testUpdatePriceShouldReturnErrorForUnauthorized() public {
        address user = address(234);
        vm.prank(user);
        marketplace.addProduct(100, 10, "c5f6f5d9-2112-4c84-8442-159747af3644");
        vm.expectRevert();
        marketplace.updatePrice(1, 0.5 ether);
    }

    function testUpdatePriceShouldUpdatePriceForEscrow() public {
        vm.prank(address(escrow));
        marketplace.addProduct(100, 10, "4991fad6-8162-4440-ba0d-d3f558fe904a");

        vm.prank(address(escrow));
        marketplace.updatePrice(1, 0.5 ether);

        (, uint256 price,) = marketplace.getProduct(1);
        assertEq(price, 0.5 ether);
    }

    function testUpdatePriceShouldUpdatePriceForSeller() public {
        marketplace.addProduct(100, 10, "f38f8a0f-92f0-4906-b4aa-dea8f8e756cb");
        marketplace.updatePrice(1, 0.5 ether);
        (, uint256 price,) = marketplace.getProduct(1);
        assertEq(price, 0.5 ether);
    }

    function testGetProductShouldReturnErrorForNonExistentProduct() public {
        vm.expectRevert();
        marketplace.getProduct(1);
    }

    function testGetProductShouldReturnProduct() public {
        marketplace.addProduct(100, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        (address seller, uint256 price, uint256 stock) = marketplace.getProduct(1);
        assertEq(seller, address(this));
        assertEq(price, 100);
        assertEq(stock, 10);
    }
}
