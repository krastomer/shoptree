package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type service struct {
	repo OrderRepository
}

var (
	ErrProductUnavailable          = errors.New("product unavailable")
	ErrDeleteProductFromCartFailed = errors.New("delete product from cart failed")
	ErrBlankCart                   = errors.New("blank cart")
	ErrWrongOwnerProduct           = errors.New("wrong owner product")
	ErrAddressNotFound             = errors.New("address not found")
	ErrUpdateOrderFailed           = errors.New("update order failed")
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

	if len(cart) == 1 {
		_ = s.repo.DeleteOrderPending(ctx, custID)
	}

	return nil
}

func (s *service) GetProductOnCart(ctx context.Context, custID int) ([]*Product, error) {
	cart, err := s.repo.GetProductPendingByCustomerID(ctx, custID)
	if err != nil {
		return nil, ErrBlankCart
	}

	var response []*Product
	for i, p := range cart {
		product, _ := s.repo.GetProductByID(ctx, p.ProductID)
		path, err := s.repo.GetImageProductByID(ctx, p.ProductID)

		response = append(response, product)
		if err == nil {
			response[i].ImagePath = path
		}
	}

	return response, nil
}

func (s *service) UpdateAddressOrder(ctx context.Context, custID, addressID int) error {
	order, err := s.repo.GetOrderPendingByCustomerID(ctx, custID)
	if err != nil {
		return ErrBlankCart
	}
	if order.CustomerID != custID {
		return ErrWrongOwnerProduct
	}

	address, err := s.repo.GetAddressesCustomer(ctx, custID)
	if err != nil {
		return ErrAddressNotFound
	}

	found := false
	for _, a := range address {
		if a.ID == addressID {
			found = true
			break
		}
	}

	if !found {
		return ErrAddressNotFound
	}

	err = s.repo.UpdateAddressOrder(ctx, order.ID, addressID)
	if err != nil {
		return ErrUpdateOrderFailed
	}
	return nil
}

func (s *service) GetCart(ctx context.Context, custID int) (*Order, error) {
	order, err := s.repo.GetOrderPendingByCustomerID(ctx, custID)
	if err != nil {
		return nil, ErrBlankCart
	}

	cart, _ := s.repo.GetProductPendingByCustomerID(ctx, custID)

	var response []*Product
	for i, p := range cart {
		product, _ := s.repo.GetProductByID(ctx, p.ProductID)
		path, err := s.repo.GetImageProductByID(ctx, p.ProductID)

		response = append(response, product)
		if err == nil {
			response[i].ImagePath = path
		}
	}

	order.Products = response
	address, err := s.repo.GetAddressCustomerByID(ctx, order.AddressID)
	if err == nil {
		order.AddressProfile = address
	}

	return order, nil
}

func (s *service) ConfirmOrder(ctx context.Context, custID int) error {
	order, err := s.repo.GetOrderPendingByCustomerID(ctx, custID)
	if err != nil {
		return ErrBlankCart
	}
	_, err = s.repo.GetProductPendingByCustomerID(ctx, custID)
	if err != nil {
		return ErrBlankCart
	}
	if order.AddressID == 0 {
		return ErrAddressNotFound
	}

	err = s.repo.UpdateStatusOrder(ctx, "VerifyPayment", order.ID)
	if err != nil {
		return ErrUpdateOrderFailed
	}

	return nil
}

func (s *service) SendPayment(ctx context.Context, c *fiber.Ctx, request *Payment, custID int) error {
	order, err := s.repo.GetOrderWaitingPaymentByCustomerID(ctx, custID)
	if err != nil {
		return ErrBlankCart
	}

	uniqueId := uuid.New()
	request.ImagePath = fmt.Sprintf("%s/%s.jpg", viper.GetString("DIRECTORY_PAYMENT"), uniqueId)
	err = c.SaveFile(request.Image, request.ImagePath)
	if err != nil {
		return ErrUpdateOrderFailed
	}

	err = s.repo.CreatePayment(ctx, order.ID, request.ImagePath)
	if err != nil {
		return ErrUpdateOrderFailed
	}

	err = s.repo.UpdateStatusOrder(ctx, "VerifyPayment", order.ID)
	if err != nil {
		return ErrUpdateOrderFailed
	}

	return nil
}
