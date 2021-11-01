package customer

import (
	"errors"
	"net/mail"
	"unicode"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type customerService struct {
	repo CustomerRepository
}

var (
	ErrPasswordGenerateBad    = errors.New("password generate bad")
	ErrRegisterCustomerFailed = errors.New("register customer failed")
	ErrEmailUsed              = errors.New("email used")
	ErrPhoneUsed              = errors.New("phone used")
	ErrPhoneInvalid           = errors.New("phone invalid")
	ErrNameInvalid            = errors.New("name invalid")
	ErrPasswordInvalid        = errors.New("password invalid")
	ErrEmailInvalid           = errors.New("email invalid")
)

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) RegisterCustomer(request *CustomerRequest) error {
	err := s.validNewCustomer(request)
	if err != nil {
		return err
	}

	cust := &Customer{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		BagLevel:    1,
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(cust.Password), viper.GetInt("BCRYPT_SIZE"))
	if err != nil {
		return ErrPasswordGenerateBad
	}

	cust.Password = string(hashPassword)

	err = s.repo.CreateCustomer(cust)
	if err != nil {
		return ErrRegisterCustomerFailed
	}
	return nil
}

func (s *customerService) validNewCustomer(cust *CustomerRequest) error {
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

	if _, err := s.repo.GetCustomerByEmail(cust.Email); err == nil {
		return ErrEmailUsed
	}

	if _, err := s.repo.GetCustomerByPhone(cust.PhoneNumber); err == nil {
		return ErrPhoneUsed
	}

	return nil
}

func (s *customerService) validPassword(password string) error {
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
