package order

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

var ()

const (
	QUERY_CREATE_NEW_ORDER = "INSERT INTO `orders` (`customer_id`, `address_id`, `status`) VALUES (?, ?, ?);"
)

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) CreatePendingProduct(ctx context.Context, product *ProductPending) (err error) {
	return err
}
