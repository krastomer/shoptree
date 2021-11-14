package customer

import (
	"context"
	"errors"
	"net/mail"
	"unicode"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo CustomerRepository
}

var (
	ErrPasswordGenerateBad         = errors.New("password generate bad")
	ErrRegisterCustomerFailed      = errors.New("register customer failed")
	ErrEmailUsed                   = errors.New("email used")
	ErrPhoneUsed                   = errors.New("phone used")
	ErrPhoneInvalid                = errors.New("phone invalid")
	ErrNameInvalid                 = errors.New("name invalid")
	ErrPasswordInvalid             = errors.New("password invalid")
	ErrEmailInvalid                = errors.New("email invalid")
	ErrCustomerNotFound            = errors.New("customer not found")
	ErrInternalServerError         = errors.New("internal server error")
	ErrAddressesCustomerNotFound   = errors.New("address customer not found")
	ErrCreateCustomerAddressFailed = errors.New("create customer address failed")
)

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &service{repo: repo}
}

func (s *service) CreateNewCustomer(ctx context.Context, request *CustomerRequest) error {
	err := s.validNewCustomer(ctx, request)
	if err != nil {
		return err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), viper.GetInt("BCRYPT_SIZE"))
	if err != nil {
		return ErrPasswordGenerateBad
	}

	request.Password = string(hashPassword)

	err = s.repo.CreateCustomer(ctx, request)
	if err != nil {
		return ErrRegisterCustomerFailed
	}
	return nil
}

func (s *service) GetAddressesCustomer(ctx context.Context, id int) ([]*Address, error) {
	address, err := s.repo.GetAddressesCustomer(ctx, id)
	if err != nil {
		return nil, ErrAddressesCustomerNotFound
	}

	return address, nil
}

func (s *service) CreateAddressCustomer(ctx context.Context, request *Address) error {
	err := s.repo.CreateAddressCustomer(ctx, request)
	if err != nil {
		return ErrCreateCustomerAddressFailed
	}
	return nil
}

func (s *service) validNewCustomer(ctx context.Context, cust *CustomerRequest) error {
	if _, err := mail.ParseAddress(cust.Email); err != nil {
		return ErrEmailInvalid
	}

	if err := s.validPassword(cust.Password); err != nil {
		return ErrPasswordInvalid
	}

	if cust.Name == "" {
		return ErrNameInvalid
	}

	if len(cust.PhoneNumber) != 10 {
		return ErrPhoneInvalid
	}

	if _, err := s.repo.GetCustomerByEmail(ctx, cust.Email); err == nil {
		return ErrEmailUsed
	}

	if _, err := s.repo.GetCustomerByPhone(ctx, cust.PhoneNumber); err == nil {
		return ErrPhoneUsed
	}

	return nil
}

func (s *service) validPassword(password string) error {
	letters := false
	number := false
	upper := false
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters = true
		case unicode.IsLetter(c) || c == ' ':
			letters = true
		}
	}
	sizeEight := len(password) >= 8
	if !(letters && number && upper && sizeEight) {
		return ErrPasswordInvalid
	}
	return nil
}
