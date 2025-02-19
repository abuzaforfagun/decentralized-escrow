package dbmodels

import "time"

type Product struct {
	Id                  string
	Seller              string
	Name                string
	Description         string
	Price               string
	Stock               string
	TransactionHash     string
	OnChainId           string
	Status              int
	CreatedAt           time.Time
	CreationCompletedAt time.Time
}

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)
