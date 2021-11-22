package order

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

const 

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) GetOrdersPending(ctx context.Context, userID int) (order *Order, err error) {
	
}
