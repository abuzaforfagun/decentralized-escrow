package models

import "math/big"

type ProductOnChain struct {
	Id     *big.Int `json:"id"`
	Seller string   `json:"seller"`
	Price  *big.Int `json:"price"`
	Qty    *big.Int `json:"qty"`
}
