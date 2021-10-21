package auth

import (
	"errors"
	"net/mail"
	"strconv"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailInvalid        = errors.New("email invalid")
	ErrPasswordInvalid     = errors.New("password invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrInternalServerError = errors.New("internal server error")
)

type authService struct {
	repo AuthRepository
}

type jwtClaims struct {
	UserID string `json:"uid"`
	User   string `json:"user"`
	jwt.StandardClaims
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(u, p string) (string, error) {
	if _, err := mail.ParseAddress(u); err != nil {
		return "", ErrEmailInvalid
	}

	if err := s.validPassword(p); err != nil {
		return "", ErrPasswordInvalid
	}

	user, err := s.findUser(u)
	if err != nil {
		return "", ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p)); err != nil {
		return "", ErrInternalServerError
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		strconv.FormatUint(uint64(user.ID), 10),
		user.Email,
		jwt.StandardClaims{
			Audience:  "shoptree-" + user.Level,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "shoptree",
		},
	})

	signedToken, err := token.SignedString([]byte("september"))

	if err != nil {
		return "", ErrInternalServerError
	}
	return signedToken, nil
}

func (s *authService) Register(*User) error {
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
	sizeEight := len(password) >= 7
	if !(letters && number && upper && sizeEight) {
		return ErrPasswordInvalid
	}
	return nil
}

func (s *authService) findUser(email string) (*User, error) {
	var cust *Customer
	var empl *Employee
	var user *User
	var err error
	c_cust := make(chan bool)
	c_empl := make(chan bool)
	defer close(c_cust)
	defer close(c_empl)
	go func() {
		cust, err = s.repo.GetCustomerByEmail(email)
		if err != nil {
			c_cust <- false
			return
		}
		c_cust <- true
	}()

	go func() {
		empl, err = s.repo.GetEmployeeByEmail(email)
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
			user = &User{
				ID:          cust.ID,
				Name:        cust.Name,
				Email:       cust.Email,
				Password:    cust.Password,
				PhoneNumber: cust.Password,
				Level:       "Customer",
			}
		case r := <-c_empl:
			if !r {
				continue
			}
			user = &User{
				ID:          empl.ID,
				Name:        empl.Name,
				Email:       empl.Email,
				Password:    empl.Password,
				PhoneNumber: cust.PhoneNumber,
				Level:       string(empl.Level),
			}
		}
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
