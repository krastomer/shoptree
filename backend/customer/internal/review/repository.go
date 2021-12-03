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
	QUERY_GET_ORDERS_DONE_WITH_REVIEW = "SELECT * FROM `orders` WHERE review IS NOT NULL AND status = 'DONE' ORDER BY created_at DESC;"
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
