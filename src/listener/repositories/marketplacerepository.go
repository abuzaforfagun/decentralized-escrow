package repositories

import (
	"context"
	"database/sql"
)

type MarketplaceRepository struct {
	db *sql.DB
}

func NewMarketplaceRepository(db *sql.DB) *MarketplaceRepository {
	return &MarketplaceRepository{
		db: db,
	}
}

func (r *MarketplaceRepository) HasProduct(ctx context.Context, correlationId string) (bool, error) {
	sql := "SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)"

	var exists bool
	err := r.db.QueryRowContext(ctx, sql, correlationId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *MarketplaceRepository) UpdateProduct(ctx context.Context, correlationId string, onchainId string, seller string, price string, stock string) error {
	sql := "UPDATE products SET on_chain_id=$1, seller=$2, price=$3, stock=$4 WHERE id=$5"
	_, err := r.db.ExecContext(ctx, sql, onchainId, seller, price, stock, correlationId)
	if err != nil {
		return err
	}
	return nil
}
