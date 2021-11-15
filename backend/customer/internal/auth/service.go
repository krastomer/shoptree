package auth

import (
	"context"
	"errors"
	"net/mail"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo AuthRepository
}

type jwtClaims struct {
	UserID int    `json:"uid"`
	User   string `json:"user"`
	jwt.StandardClaims
}

var (
	ErrEmailInvalid                = errors.New("email invalid")
	ErrEmailIncorrect              = errors.New("email incorrect")
	ErrPasswordInvalid             = errors.New("password invalid")
	ErrPasswordIncorrect           = errors.New("password incorrect")
	ErrUserNotFound                = errors.New("user not found")
	ErrTokenGenerateBad            = errors.New("token generate bad")
	ErrPasswordGenerateBad         = errors.New("password generate bad")
	ErrRegisterCustomerFailed      = errors.New("register customer failed")
	ErrEmailUsed                   = errors.New("email used")
	ErrPhoneUsed                   = errors.New("phone used")
	ErrPhoneInvalid                = errors.New("phone invalid")
	ErrNameInvalid                 = errors.New("name invalid")
	ErrCustomerNotFound            = errors.New("customer not found")
	ErrInternalServerError         = errors.New("internal server error")
	ErrAddressesCustomerNotFound   = errors.New("address customer not found")
	ErrCreateCustomerAddressFailed = errors.New("create customer address failed")
)

func NewAuthService(repo AuthRepository) AuthService {
	return &service{repo: repo}
}

func (s *service) Login(ctx context.Context, user *UserRequest) (string, error) {
	if err := s.validUserRequest(user); err != nil {
		return "", err
	}

	cust, err := s.repo.GetCustomerByEmail(ctx, user.Username)
	if err != nil {
		return "", ErrEmailIncorrect
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cust.Password), []byte(user.Password)); err != nil {
		return "", ErrPasswordIncorrect
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		cust.ID,
		cust.Email,
		jwt.StandardClaims{
			Audience:  "Customer",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "shoptree",
		},
	})

	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))

	if err != nil {
		return "", ErrTokenGenerateBad
	}
	return signedToken, nil
}

func (s *service) Register(ctx context.Context, request *Customer) error {
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

func (s *service) validUserRequest(request *UserRequest) error {
	if _, err := mail.ParseAddress(request.Username); err != nil {
		return ErrEmailInvalid
	}

	if err := s.validPassword(request.Password); err != nil {
		return err
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

func (s *service) validNewCustomer(ctx context.Context, cust *Customer) error {
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
