package services

import (
	"net/mail"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var c chan (interface{})

func init() {
	c = make(chan interface{})
}

type authService struct {
	custRepo entities.CustomerRepo
	emplRepo entities.EmployeeRepo
}

type jwtClaims struct {
	UserID string `json:"uid"`
	User   string `json:"user"`
	jwt.StandardClaims
}

func NewAuthService(custRepo entities.CustomerRepo, emplRepo entities.EmployeeRepo) entities.AuthService {
	return &authService{
		custRepo: custRepo,
		emplRepo: emplRepo,
	}
}

func (s *authService) Login(u, p string) (string, error) {
	if _, err := mail.ParseAddress(u); err != nil {
		return "", errors.ErrEmailInvalid
	}

	user, err := s.findUser(u)
	if err != nil {
		return "", errors.ErrNotFoundUser
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p)); err != nil {
		return "", errors.ErrPasswordInvlid
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
		return "", errors.ErrInternalServerError
	}

	return signedToken, nil
}

func (s *authService) findUser(email string) (*models.User, error) {
	go func() {
		cust, err := s.custRepo.GetCustomerByEmail(email)
		if err != nil {
			c <- nil
			return
		}
		c <- cust
	}()

	go func() {
		empl, err := s.emplRepo.GetEmployeeByEmail(email)
		if err != nil {
			c <- nil
			return
		}
		c <- empl
	}()

	var user interface{}
	for i := 0; i < 2; i++ {
		user = <-c
		switch t := user.(type) {
		case *models.Customer:
			return &models.User{
				ID:       t.ID,
				Name:     t.Name,
				Email:    t.Email,
				Password: t.Passwrod,
				Level:    "Customer",
			}, nil
		case *models.Employee:
			return &models.User{
				ID:       t.ID,
				Name:     t.Name,
				Email:    t.Email,
				Password: t.Password,
				Level:    string(t.Level),
			}, nil
		}
	}
	return nil, errors.ErrNotFoundUser
}
