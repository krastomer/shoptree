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
