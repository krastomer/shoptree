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
	OptsProductMR = &dbq.Options{ConcreteStruct: Product{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsOrderSR   = &dbq.Options{ConcreteStruct: Order{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_AVAILABLE_ORDER              = "SELECT * FROM `orders` WHERE orders.status = 'Pending';"
	QUERY_GET_ORDER_PENDING_BY_CUSTOMER_ID = "SELECT * FROM `orders` WHERE orders.status = 'Pending' AND orders.customer_id = ?;"
)

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) CreateOrderPending

func (r *repository) GetOrderPendingByCustomerID(ctx context.Context, customerID int) (order *Order, _ error) {
	args := []interface{}{customerID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ORDER_PENDING_BY_CUSTOMER_ID, OptsOrderSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	order = result.(*Order)
	return order, nil
}

func (r *repository) GetAvailableProduct(ctx context.Context) (products []*Product, _ error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_AVAILABLE_ORDER, OptsProductMR)
	products = result.([]*Product)
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}

	return products, nil
}
