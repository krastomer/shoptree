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
	var cust *models.Customer
	var empl *models.Employee
	var user *models.User
	var err error
	c_cust := make(chan bool)
	c_empl := make(chan bool)
	defer close(c_cust)
	defer close(c_empl)
	go func() {
		cust, err = s.custRepo.GetCustomerByEmail(email)
		if err != nil {
			c_cust <- false
			return
		}
		c_cust <- true
	}()

	go func() {
		empl, err = s.emplRepo.GetEmployeeByEmail(email)
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
			user = &models.User{
				ID:       cust.ID,
				Name:     cust.Name,
				Email:    cust.Email,
				Password: cust.Password,
				Level:    "Customer",
			}
		case r := <-c_empl:
			if !r {
				continue
			}
			user = &models.User{
				ID:       empl.ID,
				Name:     empl.Name,
				Email:    empl.Email,
				Password: empl.Password,
				Level:    string(empl.Level),
			}
		}
	}
	if user == nil {
		return nil, errors.ErrNotFoundUser
	}
	return user, nil
}

func (s *authService) Register(user *models.User) error {
	if user.Level != "Customer" {
		return errors.ErrNotAuthorized
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return errors.ErrInternalServerError
	}
	newCust := &models.Customer{
		Name:        user.Name,
		Email:       user.Email,
		Password:    string(hashPassword),
		PhoneNumber: user.Password,
	}
	err = s.custRepo.RegisterCustomer(newCust)
	return err
}
