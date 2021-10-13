package services

import (
	"net/mail"

	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type profileService struct {
	repo entities.CustomerRepo
}

func NewProfileService(repo entities.CustomerRepo) entities.ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) CreateProfile(cust *models.CustomerProfile) error {
	if _, err := mail.ParseAddress(cust.Email); err != nil {
		return errors.ErrEmailInvalid
	}
	if err := checkPasswordValid(cust.Password); err != nil {
		return errors.ErrPasswordInvalid
	}
	if len(cust.PhoneNumber) != 10 {
		return errors.ErrPhoneNumberInvalid
	}

	cust.TrimSpace()
	_, err := s.repo.GetCustomerByEmail(cust.Email)
	if err == nil {
		return errors.ErrUserExisted
	}

	_, err = s.repo.GetCustomerByPhone(cust.PhoneNumber)
	if err == nil {
		return errors.ErrPhoneNumberInvalid
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(cust.Password), 8)
	if err != nil {
		return errors.ErrInternalServerError
	}
	cust.Password = string(hashPassword)

	err = s.repo.CreateCustomer(cust)
	return err
}
