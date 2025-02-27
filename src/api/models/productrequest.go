package models

type ProductRequest struct {
	Name            string `json:"product_name"`
	Description     string `json:"product_description"`
	TransactionHash string `json:"transaction_hash"`
	CorrelationId   string `json:"correltion_id"`
}
