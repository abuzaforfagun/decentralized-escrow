package main

import (
	"database/sql"
	"fmt"

	"github.com/abuzaforfagun/decentralized-escrow/api/handlers"
	"github.com/abuzaforfagun/decentralized-escrow/api/repositories"
	"github.com/abuzaforfagun/decentralized-escrow/api/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	_ = router

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5433, "postgres", "root", "escrow")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	productRepository := repositories.NewProductRepository(db)
	onchainService := services.NewOnChainService(client, "0x5FbDB2315678afecb367f032d93F642f64180aa3")
	productService := services.NewProductService(onchainService, productRepository)
	productHandler := handlers.NewProductsHandler(productService)

	router.POST("/products", productHandler.CreateProduct)
	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
