package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/abuzaforfagun/decentralized-escrow/abigen/marketplacecontract"
	"github.com/abuzaforfagun/decentralized-escrow/listener/repositories"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type MarketplaceService struct {
	repository      *repositories.MarketplaceRepository
	client          *ethclient.Client
	contractAddress common.Address
	logger          *zap.Logger
}

func NewMarketplaceService(
	repository *repositories.MarketplaceRepository,
	client *ethclient.Client,
	contractAddress string,
	logger *zap.Logger) *MarketplaceService {
	return &MarketplaceService{
		repository:      repository,
		client:          client,
		contractAddress: common.HexToAddress(contractAddress),
		logger:          logger,
	}
}

func (s *MarketplaceService) FetchAndStoreProduct(ctx context.Context, productId *big.Int, correlationId string) error {
	marketplace, _ := marketplacecontract.NewMarketplacecontract(s.contractAddress, s.client)

	seller, price, stock, err := marketplace.GetProduct(nil, productId)
	if err != nil {
		s.logger.Error("unable to get product", zap.Error(err))
		return err
	}

	exist, err := s.repository.HasProduct(ctx, correlationId)
	if err != nil {
		s.logger.Error("unable to check does product exist", zap.Error(err))
		return err
	}
	if !exist {
		s.logger.Error("product does not exist in the off-chain database", zap.String("correlation_id", correlationId))
		return fmt.Errorf("product[%s] does not exist in the off-chain database", correlationId)
	}

	err = s.repository.UpdateProduct(ctx, correlationId, productId.String(), seller.Hex(), price.String(), stock.String())
	if err != nil {
		s.logger.Error("unable to update product", zap.Error(err))
		return err
	}

	return nil
}
