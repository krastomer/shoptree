package customer

import (
	"context"
	"errors"
)

type service struct {
	repo CustomerRepository
}

var (
	ErrAddressNotFound             = errors.New("not found address")
	ErrUserNotFound                = errors.New("user not found")
	ErrCreateAddressCustomerFailed = errors.New("create address customer failed")
	ErrDeleteAddressFailed         = errors.New("delete address failed")
)

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &service{repo: repo}
}

func (s *service) GetAddresses(ctx context.Context, custID int) ([]*Address, error) {
	addresses, err := s.repo.GetAddressesCustomer(ctx, custID)
	if err != nil {
		return nil, ErrAddressNotFound
	}

	return addresses, nil
}

func (s *service) GetCustomerProfile(ctx context.Context, custID int) (*Customer, error) {
	cust, err := s.repo.GetCustomerByID(ctx, custID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	address, _ := s.repo.GetAddressesCustomer(ctx, custID)
	cust.Address = address

	orders, err := s.repo.GetOrdersCustomer(ctx, custID)
	if err != nil {
		return cust, nil
	}

	for _, order := range orders {

		pOrder, _ := s.repo.GetProductsByOrderID(ctx, order.ID)
		for _, product := range pOrder {
			product.ImagePath, _ = s.repo.GetImageProductByID(ctx, product.ID)
		}
		orderAddress, _ := s.repo.GetAddressByOrderID(ctx, order.ID)
		order.Products = pOrder
		order.AddressProfile = orderAddress
	}
	cust.Orders = orders

	return cust, nil
}

func (s *service) CreateAddressCustomer(ctx context.Context, custID int, request *AddressRequest) error {
	address := &Address{
		CustomerID:  custID,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		AddressLine: request.AddressLine,
		Country:     request.Country,
		State:       request.State,
		City:        request.City,
		District:    request.District,
		PostalCode:  request.PostalCode,
	}
	err := s.repo.CreateAddressCustomer(ctx, address)
	if err != nil {
		return ErrCreateAddressCustomerFailed
	}
	return nil
}

func (s *service) GetAddressCustomerByID(ctx context.Context, custID int, addressID int) (*Address, error) {
	address, err := s.repo.GetAddressCustomerByID(ctx, addressID)
	if err != nil {
		return nil, ErrAddressNotFound
	}

	if address.CustomerID != custID {
		return nil, ErrAddressNotFound
	}

	return address, nil
}

func (s *service) DeleteAddressCustomer(ctx context.Context, custID int, addressID int) error {
	address, err := s.repo.GetAddressCustomerByID(ctx, addressID)
	if err != nil {
		return ErrAddressNotFound
	}

	if address.CustomerID != custID {
		return ErrAddressNotFound
	}

	err = s.repo.DeleteAddressCustomer(ctx, addressID)
	if err != nil {
		return ErrDeleteAddressFailed
	}

	return nil
}
