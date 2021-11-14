package order

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type repository struct {
	db  *sql.DB
	rdb *redis.Client
}

var ()

const (
	QUERY_CREATE_NEW_ORDER = "INSERT INTO `orders` (`customer_id`, `address_id`, `status`) VALUES (?, ?, ?);"
)

func NewOrderRepository(db *sql.DB, rds *redis.Client) OrderRepository {
	return &repository{db: db, rdb: rds}
}

func (r *repository) CreatePendingProduct(ctx context.Context, product *ProductPending) (err error) {
	return err
}
