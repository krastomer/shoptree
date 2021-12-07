package review

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
	QUERY_GET_ORDERS_DONE_WITH_REVIEW = "SELECT * FROM `orders` WHERE review != 'NULL' AND status = 'DONE' ORDER BY created_at DESC;"
	QUERY_GET_ORDERS_DONE_CUSTOMER    = "SELECT * FROM `orders` WHERE review = 'NULL' AND status = 'DONE' AND customer_id = ?;"
	QUERY_UPDATE_ORDER_REVIEW         = "UPDATE `orders` SET `review` = ? WHERE `orders`.`id` = ?;"
)

func NewReviewRepository(db *sql.DB) ReviewRepository {
	return &repository{db: db}
}

func (r *repository) GetOrdersDoneWithReview(ctx context.Context) (orders []*Order, _ error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_ORDERS_DONE_WITH_REVIEW, OptsOrderMR)
	orders = result.([]*Order)
	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	return orders, nil
}

func (r *repository) GetOrdersDoneCustomer(ctx context.Context, custID int) (orders []*Order, _ error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_ORDERS_DONE_CUSTOMER, OptsOrderMR, custID)
	orders = result.([]*Order)
	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	return orders, nil
}

func (r *repository) UpdateOrderReview(ctx context.Context, orderID int, review string) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_UPDATE_ORDER_REVIEW, nil, review, orderID)
		if err != nil {
			return
		}
		txCommit()
	})

	return err
}
