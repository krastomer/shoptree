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
	ErrEmailInvalid      = errors.New("email invalid")
	ErrEmailIncorrect    = errors.New("email incorrect")
	ErrPasswordInvalid   = errors.New("password invalid")
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrUserNotFound      = errors.New("user not found")
	ErrTokenGenerateBad  = errors.New("token generate bad")
)

func NewAuthService(repo AuthRepository) AuthService {
	return &service{repo: repo}
}

func (s *service) Login(ctx context.Context, user *UserRequest) (string, error) {
	if err := s.validUserRequest(user); err != nil {
		return "", err
	}

	empl, err := s.repo.GetEmployeeByEmail(ctx, user.Username)
	if err != nil {
		return "", ErrEmailIncorrect
	}

	if err := bcrypt.CompareHashAndPassword([]byte(empl.Password), []byte(user.Password)); err != nil {
		return "", ErrPasswordIncorrect
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		empl.ID,
		empl.Email,
		jwt.StandardClaims{
			Audience:  empl.Level,
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
