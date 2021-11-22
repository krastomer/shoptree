package order

import (
	"context"
	"errors"
)

type service struct {
	repo OrderRepository
}

var (
	ErrProductUnavailable = errors.New("product unavailable")
)

func NewOrderService(repo OrderRepository) OrderService {
	return &service{repo: repo}
}

func (s *service) AddProductToCart(ctx context.Context, custID, prodID int) error {
	_, err := s.repo.GetAvailableProductByID(ctx, prodID)
	if err != nil {
		return ErrProductUnavailable
	}

	orderID, err := s.repo.GetOrderPendingByCustomerID(ctx, custID)
	if err != nil {
		_ = s.repo.CreateOrderPending(ctx, orderID.ID)
	}

	return nil

}
