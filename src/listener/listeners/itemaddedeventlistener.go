package listeners

import (
	"context"

	"github.com/abuzaforfagun/decentralized-escrow/abigen/marketplacecontract"
	"github.com/abuzaforfagun/decentralized-escrow/listener/services"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type ItemAddedEventListener struct {
	client          *ethclient.Client
	contractAddress common.Address
	logger          *zap.Logger
	service         *services.MarketplaceService
}

func NewItemAddedEventListener(client *ethclient.Client, service *services.MarketplaceService, contractAddress string, logger *zap.Logger) *ItemAddedEventListener {
	return &ItemAddedEventListener{
		client:          client,
		contractAddress: common.HexToAddress(contractAddress),
		logger:          logger,
		service:         service,
	}
}

func (l *ItemAddedEventListener) StartListening(ctx context.Context) (chan error, error) {

	marketplace, err := marketplacecontract.NewMarketplacecontractFilterer(l.contractAddress, l.client)
	if err != nil {
		l.logger.Error("unable to create filterer", zap.Error(err))
		return nil, err
	}

	itemAddedCh := make(chan *marketplacecontract.MarketplacecontractItemAdded)
	subscription, err := marketplace.WatchItemAdded(&bind.WatchOpts{Context: ctx}, itemAddedCh, nil, nil)
	if err != nil {
		l.logger.Error("unable to create watch item added subscription", zap.Error(err))
		return nil, err
	}

	go func() {
		for event := range itemAddedCh {
			for retried := 1; retried <= 3; retried++ {
				err := l.service.FetchAndStoreProduct(ctx, event.Id, event.Guid)
				if err == nil {
					break
				}
				l.logger.Error("unable to fetch and store product", zap.Int("retried", retried), zap.Error(err))
			}
		}
	}()

	<-subscription.Err()
	return nil, err
}
