package services

import (
	"context"
	"fmt"

	"github.com/abuzaforfagun/decentralized-escrow/api/models"
	"github.com/abuzaforfagun/decentralized-escrow/api/repositories"
)

type ProductService struct {
	onChainService    *onChainService
	productRepository *repositories.ProductRepository
}

func NewProductService(onChainService *onChainService, productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{onChainService: onChainService, productRepository: productRepository}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.ProductRequest) error {
	isPending, err := s.onChainService.isPendingTransaction(ctx, product.TransactionHash)

	if err != nil {
		return fmt.Errorf("failed to check transaction status: %w", err)
	}

	if isPending {
		err := s.productRepository.AddDraftProduct(ctx, *product)
		if err == nil {
			return fmt.Errorf("failed to save product details: %w", err)
		}
		return nil
	}

	productDetails, err := s.onChainService.getProductDetailsByTransactionHash(ctx, product.TransactionHash)
	if err != nil {
		return fmt.Errorf("failed to get product details: %w", err)
	}

	_ = productDetails
	err = s.productRepository.AddProduct(ctx, productDetails.Id.String(), product.Name, product.Description, product.TransactionHash, productDetails.Seller, productDetails.Price.String(), productDetails.Qty.String())
	if err != nil {
		return fmt.Errorf("failed to save product details: %w", err)
	}
	return nil
}
