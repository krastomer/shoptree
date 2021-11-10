package order

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/rocketlaunchr/dbq/v2"
)

type mariaDBRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

var ()

const (
	QUERY_CREATE_NEW_ORDER = "INSERT INTO `orders` (`customer_id`, `address_id`, `status`) VALUES (?, ?, ?);"
)

func NewOrderRepository(db *sql.DB, rdb *redis.Client) OrderRepository {
	return &mariaDBRepository{db: db, rdb: rdb}
}

func (r *mariaDBRepository) CreateNewOrder(ctx context.Context, order *Order) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_CREATE_NEW_ORDER, nil,
			order.CustomerID,
			order.AddressID,
			order.Status,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}
