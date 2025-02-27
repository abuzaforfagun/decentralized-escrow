// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Test} from "forge-std/Test.sol";
import {Marketplace} from "../src/Marketplace.sol";
import {Escrow} from "../src/Escrow.sol";
import {DeployContracts} from "../script/DeployContracts.s.sol";

contract EscrowTest is Test {
    Escrow escrow;
    Marketplace marketplace;
    DeployContracts deployer;

    function setUp() public {
        deployer = new DeployContracts();
        (marketplace, escrow) = deployer.run();
    }

    function testCreateEscrowShouldCreateEscrow() public {
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);
        assertEq(escrowId, 1);
    }

    function testCreateEscrowShouldSetStatusToPending() public {
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);
        assertEq(escrowId, 1);

        (,,,, uint8 status) = escrow.getEscrow(escrowId);
        assertEq(uint256(status), uint256(Escrow.EscrowStatus.Pending));
    }

    function testCreateEscrowShouldReturnErrorForInvalidProduct() public {
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        vm.expectRevert();
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);
    }

    function testCreateEscrowShouldUpdateProductStock() public {
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.prank(address(123));
        vm.deal(address(123), 100 ether);
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        (,, uint256 stock) = marketplace.getProduct(productId);
        assertEq(stock, 5);
    }

    function testCreateEscrowShouldReturnErrorForInefficiantEth() public {
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.prank(address(123));
        vm.deal(address(123), 50 ether);
        vm.expectRevert();
        escrow.createEscrow{value: 4 ether}(productId, 5);

        (,, uint256 stock) = marketplace.getProduct(productId);
        assertEq(stock, 10);
    }

    function testCreateEscrowShouldReturnErrorForStockNotAvailable() public {
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.prank(address(123));
        vm.deal(address(123), 100 ether);
        vm.expectRevert();
        escrow.createEscrow{value: 15 ether}(productId, 15);

        (,, uint256 stock) = marketplace.getProduct(productId);
        assertEq(stock, 10);
    }

    function testReceivalConfirmationShouldCompleteEscrow() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(address(123));
        escrow.receivalConfirmation(escrowId);
        (,,,, uint8 status) = escrow.getEscrow(escrowId);
        assertEq(uint256(status), uint256(Escrow.EscrowStatus.Completed));
    }

    function testReceivalConfirmationShouldReturnErrorWhenBuyerDoesNotCall() public {
        marketplace.addProduct(100, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.prank(address(123));
        vm.deal(address(123), 100 ether);
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);
        vm.prank(address(234));
        vm.expectRevert();
        escrow.receivalConfirmation(escrowId);
    }

    function testReceivalConfirmationShouldReturnErrorWhenEscrowIsInCompletedStatus() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(address(123));
        escrow.receivalConfirmation(escrowId);

        vm.expectRevert();
        escrow.receivalConfirmation(escrowId);
    }

    function testReceivalConfirmationShouldReturnErrorWhenEscrowIsInResolvedStatus() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);
        vm.prank(msg.sender);
        escrow.resolve(escrowId);

        vm.expectRevert();
        vm.prank(address(123));
        escrow.receivalConfirmation(escrowId);
    }

    function testReceivalConfirmationShouldReturnErrorWhenEscrowIsInRefundStatus() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);
        vm.prank(msg.sender);
        escrow.refund(escrowId);

        vm.expectRevert();
        vm.prank(address(123));
        escrow.receivalConfirmation(escrowId);
    }

    function testReceivalConfirmationShouldSendMoneyToSeller() public {
        address seller = address(234);
        vm.prank(seller);
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        uint256 balanceBeforeConfirmation = seller.balance;

        vm.prank(address(123));
        escrow.receivalConfirmation(escrowId);

        uint256 balanceAfterConfirmation = seller.balance;
        assertEq(balanceAfterConfirmation, balanceBeforeConfirmation + 5 ether);
    }

    function testDisputeShouldSetStatusToDisputed() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        (,,,, uint8 status) = escrow.getEscrow(escrowId);
        assertEq(uint256(status), uint256(Escrow.EscrowStatus.Disputed));
    }

    function testDisputeShouldRevertIfNotOwner() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(address(123));
        vm.expectRevert();
        escrow.dispute(escrowId);
    }

    function testDisputeShouldRevertIfNotPending() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");
        uint256 productId = 1;
        vm.deal(address(123), 100 ether);
        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(address(123));
        escrow.receivalConfirmation(escrowId);

        vm.expectRevert();
        vm.prank(address(this));
        escrow.dispute(escrowId);
    }

    function testResolveShouldSetStatusToResolvedAndTransferFunds() public {
        address seller = address(234);

        vm.prank(seller);
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(address(123), 100 ether);

        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        uint256 balanceBeforeResolution = seller.balance;

        vm.prank(msg.sender);
        escrow.resolve(escrowId);

        (,,,, uint8 status) = escrow.getEscrow(escrowId);
        assertEq(uint256(status), uint256(Escrow.EscrowStatus.Resolved));

        uint256 balanceAfterResolution = seller.balance;
        assertEq(balanceAfterResolution, balanceBeforeResolution + 5 ether);
    }

    function testResolveShouldRevertIfNotOwner() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(address(123), 100 ether);

        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        vm.prank(address(123));
        vm.expectRevert();
        escrow.resolve(escrowId);
    }

    function testResolveShouldRevertIfNotDisputed() public {
        vm.prank(address(234));
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(address(123), 100 ether);

        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(address(this));
        vm.expectRevert();
        escrow.resolve(escrowId);
    }

    function testResolveShouldRevertIfTransferFails() public {
        address seller = address(234);

        vm.prank(seller);
        marketplace.addProduct(1 ether, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(address(123), 100 ether);

        vm.prank(address(123));
        uint256 escrowId = escrow.createEscrow{value: 5 ether}(productId, 5);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        vm.etch(seller, hex"");

        vm.prank(address(this));
        vm.expectRevert();
        escrow.resolve(escrowId);
    }

    function testRefundShouldSetStatusToRefundedAndTransferFunds() public {
        address buyer = address(123);
        address seller = address(234);
        uint256 price = 2 ether;
        uint256 qty = 3;
        uint256 totalAmount = price * qty;

        vm.prank(seller);
        marketplace.addProduct(price, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(buyer, 100 ether);

        vm.prank(buyer);
        uint256 escrowId = escrow.createEscrow{value: totalAmount}(productId, qty);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        uint256 balanceBeforeRefund = buyer.balance;

        vm.prank(msg.sender);
        escrow.refund(escrowId);

        (,,,, uint8 status) = escrow.getEscrow(escrowId);
        assertEq(uint256(status), uint256(Escrow.EscrowStatus.Refunded));

        uint256 balanceAfterRefund = buyer.balance;
        assertEq(balanceAfterRefund, balanceBeforeRefund + totalAmount);
    }

    function testRefundBySellerShouldSetStatusToRefundedAndTransferFunds() public {
        address buyer = address(123);
        address seller = address(234);
        uint256 price = 2 ether;
        uint256 qty = 3;
        uint256 totalAmount = price * qty;

        vm.prank(seller);
        marketplace.addProduct(price, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(buyer, 100 ether);

        vm.prank(buyer);
        uint256 escrowId = escrow.createEscrow{value: totalAmount}(productId, qty);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        uint256 balanceBeforeRefund = buyer.balance;

        vm.prank(seller);
        escrow.refund(escrowId);

        (,,,, uint8 status) = escrow.getEscrow(escrowId);
        assertEq(uint256(status), uint256(Escrow.EscrowStatus.Refunded));

        uint256 balanceAfterRefund = buyer.balance;
        assertEq(balanceAfterRefund, balanceBeforeRefund + totalAmount);
    }

    function testRefundShouldRevertIfNotOwnerOrSeller() public {
        address buyer = address(123);
        address seller = address(234);
        address nonOwner = address(999);
        uint256 price = 2 ether;
        uint256 qty = 3;
        uint256 totalAmount = price * qty;

        vm.prank(seller);
        marketplace.addProduct(price, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(buyer, 100 ether);

        vm.prank(buyer);
        uint256 escrowId = escrow.createEscrow{value: totalAmount}(productId, qty);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        vm.prank(nonOwner);
        vm.expectRevert();
        escrow.refund(escrowId);
    }

    function testRefundShouldRevertIfStatusNotDisputedOrPending() public {
        address buyer = address(123);
        address seller = address(234);
        uint256 price = 2 ether;
        uint256 qty = 3;
        uint256 totalAmount = price * qty;

        vm.prank(seller);
        marketplace.addProduct(price, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(buyer, 100 ether);

        vm.prank(buyer);
        uint256 escrowId = escrow.createEscrow{value: totalAmount}(productId, qty);

        vm.prank(buyer);
        escrow.receivalConfirmation(escrowId);

        vm.prank(msg.sender);
        vm.expectRevert();
        escrow.refund(escrowId);
    }

    function testRefundShouldRevertIfTransferFails() public {
        address buyer = address(123);
        address seller = address(234);
        uint256 price = 2 ether;
        uint256 qty = 3;
        uint256 totalAmount = price * qty;

        vm.prank(seller);
        marketplace.addProduct(price, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(buyer, 100 ether);

        vm.prank(buyer);
        uint256 escrowId = escrow.createEscrow{value: totalAmount}(productId, qty);

        vm.prank(msg.sender);
        escrow.dispute(escrowId);

        vm.etch(buyer, hex"");

        vm.prank(address(this));
        vm.expectRevert();
        escrow.refund(escrowId);
    }

    function testGetEscrowShouldReturnCorrectEscrowData() public {
        address buyer = address(123);
        address seller = address(234);
        uint256 price = 2 ether;
        uint256 qty = 3;
        uint256 totalAmount = price * qty;

        vm.prank(seller);
        marketplace.addProduct(price, 10, "fa306437-0e41-48b8-b42b-98f07bd783b7");

        uint256 productId = 1;
        vm.deal(buyer, 100 ether);

        vm.prank(buyer);
        uint256 escrowId = escrow.createEscrow{value: totalAmount}(productId, qty);

        (address escrowBuyer, address escrowSeller, uint256 escrowPrice, uint256 escrowQty, uint8 escrowStatus) =
            escrow.getEscrow(escrowId);

        assertEq(escrowBuyer, buyer);
        assertEq(escrowSeller, seller);
        assertEq(escrowPrice, price);
        assertEq(escrowQty, qty);
        assertEq(escrowStatus, uint8(Escrow.EscrowStatus.Pending));
    }

    function testGetEscrowShouldRevertForInvalidEscrowId() public {
        uint256 invalidEscrowId = 9999;

        vm.expectRevert();
        escrow.getEscrow(invalidEscrowId);
    }
}
