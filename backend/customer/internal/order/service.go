package order

import (
	"context"
	"errors"
)

type service struct {
	repo OrderRepository
}

var (
	ErrProductUnavailable          = errors.New("product unavailable")
	ErrDeleteProductFromCartFailed = errors.New("delete product from cart failed")
	ErrBlankCart                   = errors.New("blank cart")
	ErrWrongOwnerProduct           = errors.New("wrong owner product")
)

func NewOrderService(repo OrderRepository) OrderService {
	return &service{repo: repo}
}

func (s *service) AddProductToCart(ctx context.Context, custID int, prodID int) error {
	_, err := s.repo.GetAvailableProductByID(ctx, prodID)
	if err != nil {
		return ErrProductUnavailable
	}

	order, err := s.repo.GetOrderPendingByCustomerID(ctx, custID)
	if err != nil {
		_ = s.repo.CreateOrderPending(ctx, custID)
		order, _ = s.repo.GetOrderPendingByCustomerID(ctx, custID)
	}

	err = s.repo.AddProductToOrder(ctx, order.ID, prodID)

	return err
}

func (s *service) RemoveProductFromCart(ctx context.Context, custID int, prodID int) error {
	cart, err := s.repo.GetProductPendingByCustomerID(ctx, custID)
	if err != nil {
		return ErrBlankCart
	}

	found := false
	for _, product := range cart {
		if product.ProductID == prodID {
			found = true
			break
		}
	}

	if !found {
		return ErrWrongOwnerProduct
	}

	err = s.repo.DeleteProductFromOrder(ctx, prodID)
	if err != nil {
		return ErrDeleteProductFromCartFailed
	}

	return nil
}

func (s *service) GetProductOnCart(ctx context.Context, custID int) ([]*Product, error) {
	cart, err := s.repo.GetProductPendingByCustomerID(ctx, custID)
	if err != nil {
		return nil, ErrBlankCart
	}

	var response []*Product
	for _, p := range cart {
		product, _ := s.repo.GetProductByID(ctx, p.ProductID)
		response = append(response, product)
	}

	return response, nil
}
