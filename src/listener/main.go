package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/abuzaforfagun/decentralized-escrow/listener/listeners"
	"github.com/abuzaforfagun/decentralized-escrow/listener/repositories"
	"github.com/abuzaforfagun/decentralized-escrow/listener/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var client *ethclient.Client
var logger *zap.Logger

var ContractAddress string

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panic("unable to init zap")
	}

	defer logger.Sync()

	err = godotenv.Load(".env")
	if err != nil {
		logger.Error("unable to load .env file", zap.Error(err))
	}

	nodeWs := os.Getenv("NODE_WS")
	client, err = ethclient.Dial(nodeWs)
	if err != nil {
		logger.Error("unable to connect with node", zap.Error(err))
	}

	host, portStr, user, password, dbname := os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Panic("invalid port", zap.Error(err))
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Panic("unable to connect database", zap.Error(err))
	}

	marketplaceContractAddress := os.Getenv("MARKETPLACE_CONTRACT_ADDRESS")
	if marketplaceContractAddress == "" {
		logger.Panic("unable to get marketplace contract address from enviornment")
	}

	repository := repositories.NewMarketplaceRepository(db)

	marketplaceService := services.NewMarketplaceService(repository, client, marketplaceContractAddress, logger)

	itemAddedEventListener := listeners.NewItemAddedEventListener(client, marketplaceService, marketplaceContractAddress, logger)

	logger.Info("listener is starting...")
	errCh, err := itemAddedEventListener.StartListening(context.Background())
	if err != nil {
		logger.Panic("unable to start listening", zap.Error(err))
	}

	select {
	case err := <-errCh:
		logger.Panic("program close due to error", zap.Error(err))
	}
}
