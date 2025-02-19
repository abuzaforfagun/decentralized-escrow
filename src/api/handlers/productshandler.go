package handlers

import (
	"github.com/abuzaforfagun/decentralized-escrow/api/models"
	"github.com/abuzaforfagun/decentralized-escrow/api/services"
	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	productService *services.ProductService
}

func NewProductsHandler(productService *services.ProductService) *ProductsHandler {
	return &ProductsHandler{productService: productService}
}

func (h *ProductsHandler) CreateProduct(c *gin.Context) {
	var requestModel models.ProductRequest

	err := c.BindJSON(&requestModel)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err = h.productService.CreateProduct(c, &requestModel)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
}
