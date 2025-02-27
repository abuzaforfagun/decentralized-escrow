package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/abuzaforfagun/decentralized-escrow/api/dbmodels"
	"github.com/abuzaforfagun/decentralized-escrow/api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) AddDraftProduct(ctx context.Context, product models.ProductRequest) error {
	sql := "INSERT INTO products (id, name, description, transaction_hash, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.ExecContext(ctx, sql, product.CorrelationId, product.Name, product.Description, product.TransactionHash, dbmodels.Pending, time.Now().UTC())

	return err
}

func (r *ProductRepository) AddProduct(ctx context.Context, id string, onchainId string, name string, description string, transactionHash string, seller string, price string, qty string) error {
	sql := "INSERT INTO products (id, on_chain_id, name, description, transaction_hash, seller, price, stock, status, created_at, creation_completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	res, err := r.db.ExecContext(ctx, sql, id, onchainId, name, description, transactionHash, seller, price, qty, dbmodels.Approved, time.Now().UTC(), time.Now().UTC())
	_ = res
	return err
}
