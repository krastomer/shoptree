package order

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsOrderMR = &dbq.Options{ConcreteStruct: Order{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_CREATE_NEW_ORDER   = "INSERT INTO `orders` (`customer_id`, `address_id`, `status`) VALUES (?, ?, ?);"
	QUERY_GET_PENDING_ORDERS = "SELECT * FROM `orders` WHERE created_at > NOW() - INTERVAL 1 HOUR;"
)

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) GetPendingOrders(ctx context.Context) (orders []*Order, err error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_PENDING_ORDERS, OptsOrderMR)
	orders = result.([]*Order)
	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	return orders, nil
}

func (r *repository) CreatePendingOrder(ctx context.Context, product *OrderPending) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_CREATE_NEW_ORDER, nil,
			product.CustomerID,
			nil,
			product.Status,
			product.ExpiresAt,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}
