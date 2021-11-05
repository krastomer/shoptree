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

type authService struct {
	repo AuthRepository
}

type jwtClaims struct {
	UserID int    `json:"uid"`
	User   string `json:"user"`
	jwt.StandardClaims
}

var (
	ErrEmailInvalid      = errors.New("email invalid")
	ErrEmailIncorrect    = errors.New("email incorrect")
	ErrPasswordInvalid   = errors.New("password invalid")
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrUserNotFound      = errors.New("user not found")
	ErrTokenGenerateBad  = errors.New("token generate bad")
)

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(ctx context.Context, user *UserRequest) (string, error) {
	if err := s.validUserRequest(user); err != nil {
		return "", err
	}

	userToken, err := s.findUser(ctx, user.Username)
	if err != nil {
		return "", ErrEmailIncorrect
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userToken.Password), []byte(user.Password)); err != nil {
		return "", ErrPasswordIncorrect
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		userToken.ID,
		userToken.Email,
		jwt.StandardClaims{
			Audience:  userToken.Level,
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

func (s *authService) findUser(ctx context.Context, email string) (*UserToken, error) {
	var cust *Customer
	var empl *Employee
	var user *UserToken
	var err error
	c_cust := make(chan bool)
	c_empl := make(chan bool)
	defer close(c_cust)
	defer close(c_empl)
	go func() {
		cust, err = s.repo.GetCustomerByEmail(ctx, email)
		if err != nil {
			c_cust <- false
			return
		}
		c_cust <- true
	}()

	go func() {
		empl, err = s.repo.GetEmployeeByEmail(ctx, email)
		if err != nil {
			c_empl <- false
			return
		}
		c_empl <- true
	}()

	for i := 0; i < 2; i++ {
		select {
		case r := <-c_cust:
			if !r {
				continue
			}
			user = &UserToken{
				ID:       cust.ID,
				Email:    cust.Email,
				Password: cust.Password,
				Level:    "Customer",
			}
		case r := <-c_empl:
			if !r {
				continue
			}
			user = &UserToken{
				ID:       empl.ID,
				Email:    empl.Email,
				Password: empl.Password,
				Level:    string(empl.Level),
			}
		}
	}

	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *authService) validUserRequest(request *UserRequest) error {

	if _, err := mail.ParseAddress(request.Username); err != nil {
		return ErrEmailInvalid
	}

	if err := s.validPassword(request.Password); err != nil {
		return err
	}

	return nil
}

func (s *authService) validPassword(password string) error {
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
