package services

import (
	"context"
	"math/big"
	"strings"

	"github.com/abuzaforfagun/decentralized-escrow/api/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type onChainService struct {
	client  *ethclient.Client
	address *common.Address
}

func NewOnChainService(client *ethclient.Client, address string) *onChainService {
	addressObj := common.HexToAddress(address)
	return &onChainService{client: client, address: &addressObj}
}

func (s *onChainService) isPendingTransaction(ctx context.Context, transactionHash string) (bool, error) {
	_, isPending, err := s.client.TransactionByHash(ctx, common.HexToHash(transactionHash))
	return isPending, err
}

func (s *onChainService) getProductDetailsByTransactionHash(ctx context.Context, transactionHash string) (*models.ProductOnChain, error) {
	receipt, err := s.client.TransactionReceipt(ctx, common.HexToHash(transactionHash))

	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(`[
  {
      "type": "event",
      "name": "ItemAdded",
      "inputs": [
        {
          "name": "id",
          "type": "uint256",
          "indexed": true,
          "internalType": "uint256"
        },
        {
          "name": "seller",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "guid",
          "type": "string",
          "indexed": false,
          "internalType": "string"
        }
      ],
      "anonymous": false
    }
]`))
	if err != nil {
		return nil, err
	}
	eventSignature := contractABI.Events["ItemAdded"].ID
	var productId *big.Int
	for _, log := range receipt.Logs {
		if log.Topics[0] == eventSignature {
			event := struct {
				ProductID *big.Int
				Seller    common.Address
				Guid      string
			}{}

			err := contractABI.UnpackIntoInterface(&event, "ItemAdded", log.Data)
			if err != nil {
				return nil, err
			}

			productId = new(big.Int).SetBytes(log.Topics[1].Bytes())
		}
	}
	productDetails, err := s.getProductById(ctx, productId)
	if err != nil {
		return nil, err
	}

	return productDetails, nil
}

func (s *onChainService) getProductById(ctx context.Context, productId *big.Int) (*models.ProductOnChain, error) {
	contractABI, err := abi.JSON(strings.NewReader(`[{
    "type": "function",
    "name": "getProduct",
    "inputs": [
      {
        "name": "id",
        "type": "uint256",
        "internalType": "uint256"
      }
    ],
    "outputs": [
      {
        "name": "",
        "type": "address",
        "internalType": "address"
      },
      {
        "name": "",
        "type": "uint256",
        "internalType": "uint256"
      },
      {
        "name": "",
        "type": "uint256",
        "internalType": "uint256"
      }
    ],
    "stateMutability": "view"
  }]`))
	if err != nil {
		return nil, err
	}

	callData, err := contractABI.Pack("getProduct", productId)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   s.address,
		Data: callData,
	}

	result, err := s.client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	var seller common.Address
	var price, stock *big.Int

	err = contractABI.UnpackIntoInterface(&[]interface{}{&seller, &price, &stock}, "getProduct", result)
	if err != nil {
		return nil, err
	}

	productDetails := &models.ProductOnChain{
		Id:     productId,
		Seller: seller.Hex(),
		Price:  price,
		Qty:    stock,
	}
	return productDetails, nil
}
