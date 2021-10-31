package auth

import (
	"errors"
	"time"

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
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrUserNotFound      = errors.New("user not found")
	ErrTokenGenerate     = errors.New("token generate bad")
)

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(user *UserRequest) (string, error) {
	userToken, err := s.findUser(user.Username)
	if err != nil {
		return "", err
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
		return "", ErrTokenGenerate
	}
	return signedToken, nil
}

func (s *authService) findUser(email string) (*UserToken, error) {
	var cust *Customer
	var empl *Employee
	var user *UserToken
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
