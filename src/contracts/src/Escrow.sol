// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Marketplace} from "./Marketplace.sol";
import {Models} from "./libraries/Product.sol";

contract Escrow {
    error Escrow_InefficiantEth();
    error Escrow_InvalidRequest();
    error Escrow_Unauthorized();
    error Escrow_Failed();
    error Escrow_StockNotAvailable();

    event EscrowCreated(uint256 indexed escrowId);
    event EscrowCompleted(uint256 indexed escrowId);
    event EscrowDisputed(uint256 indexed escrowId);
    event EscrowResolved(uint256 indexed escrowId);
    event EscrowRefunded(uint256 indexed escrowId);

    enum EscrowStatus {
        Pending,
        Completed,
        Disputed,
        Resolved,
        Refunded
    }

    struct EscrowData {
        uint256 itemId;
        address buyer;
        address seller;
        uint256 price;
        uint256 qty;
        EscrowStatus status;
    }

    mapping(uint256 => EscrowData) private s_escrow;

    uint256 private s_escrowCounter;
    Marketplace private immutable i_marketplace;
    address private immutable i_owner;

    modifier onlyOwnerOrSeller(uint256 escrowId) {
        if (msg.sender != i_owner && msg.sender != s_escrow[escrowId].seller) {
            revert Escrow_Unauthorized();
        }
        _;
    }

    modifier onlyBuyer(uint256 escrowId) {
        if (msg.sender != s_escrow[escrowId].buyer) {
            revert Escrow_Unauthorized();
        }
        _;
    }

    modifier onlySeller(uint256 escrowId) {
        if (msg.sender != s_escrow[escrowId].seller) {
            revert Escrow_Unauthorized();
        }
        _;
    }

    modifier onlyOwner() {
        if (msg.sender != i_owner) {
            revert Escrow_Unauthorized();
        }
        _;
    }

    constructor(address _marketplace) {
        i_marketplace = Marketplace(_marketplace);
        i_owner = msg.sender;
    }

    function createEscrow(uint256 productId, uint256 qty) external payable returns (uint256) {
        (address seller, uint256 price, uint256 stock) = i_marketplace.getProduct(productId);

        if (msg.value < (price * qty)) {
            revert Escrow_InefficiantEth();
        }

        if (stock < qty) {
            revert Escrow_StockNotAvailable();
        }
        s_escrowCounter++;
        uint256 escrowId = s_escrowCounter;

        s_escrow[escrowId] = EscrowData({
            itemId: productId,
            seller: seller,
            buyer: msg.sender,
            price: price,
            qty: qty,
            status: EscrowStatus.Pending
        });
        uint256 newStock = stock - qty;
        i_marketplace.updateStock(productId, newStock);
        emit EscrowCreated(s_escrowCounter);
        return escrowId;
    }

    function receivalConfirmation(uint256 escrowId) external onlyBuyer(escrowId) {
        EscrowData storage escrow = s_escrow[escrowId];
        if (escrow.status != EscrowStatus.Pending && escrow.status != EscrowStatus.Disputed) {
            revert Escrow_InvalidRequest();
        }

        (bool success,) = escrow.seller.call{value: escrow.price * escrow.qty}("");
        if (!success) {
            revert Escrow_Failed();
        }

        escrow.status = EscrowStatus.Completed;
        emit EscrowCompleted(escrowId);
    }

    function dispute(uint256 escrowId) external onlyOwner {
        EscrowData storage escrow = s_escrow[escrowId];

        if (escrow.status != EscrowStatus.Pending) {
            revert Escrow_InvalidRequest();
        }

        escrow.status = EscrowStatus.Disputed;
        emit EscrowDisputed(escrowId);
    }

    function resolve(uint256 escrowId) external onlyOwner {
        EscrowData storage escrow = s_escrow[escrowId];

        if (escrow.status != EscrowStatus.Disputed) {
            revert Escrow_InvalidRequest();
        }

        (bool success,) = escrow.seller.call{value: escrow.price * escrow.qty}("");
        if (!success) {
            revert Escrow_Failed();
        }
        escrow.status = EscrowStatus.Resolved;
        emit EscrowResolved(escrowId);
    }

    function refund(uint256 escrowId) external onlyOwnerOrSeller(escrowId) {
        EscrowData storage escrow = s_escrow[escrowId];

        if (escrow.status != EscrowStatus.Disputed && escrow.status != EscrowStatus.Pending) {
            revert Escrow_InvalidRequest();
        }

        (bool success,) = escrow.buyer.call{value: escrow.price * escrow.qty}("");
        if (!success) {
            revert Escrow_Failed();
        }
        escrow.status = EscrowStatus.Refunded;
        emit EscrowRefunded(escrowId);
    }

    function getEscrow(uint256 escrowId) external view returns (address, address, uint256, uint256, uint8) {
        EscrowData memory escrow = s_escrow[escrowId];

        if (escrow.buyer == address(0)) {
            revert Escrow_InvalidRequest();
        }
        return (escrow.buyer, escrow.seller, escrow.price, escrow.qty, uint8(escrow.status));
    }
}
