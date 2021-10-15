package services

import (
	"sync"

	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
)

type profileService struct {
	repo entities.CustomerProfileRepo
}

func NewProfileService(repo entities.CustomerProfileRepo) entities.ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) GetProfile(id uint32) (*models.CustomerProfile, error) {
	custPro := &models.CustomerProfile{}

	c := make(chan bool)
	defer close(c)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cust, err := s.repo.GetCustomerByID(id)
		if err != nil {
			c <- false
			return
		}
		custPro.Name = cust.Name
		custPro.Email = cust.Email
		custPro.PhoneNumber = cust.PhoneNumber
		custPro.CreatedAt = cust.CreatedAt
		c <- true
	}()

	go func() {
		defer wg.Done()
		addresses, err := s.repo.GetAddresses(id)
		if err != nil {
			return
		}
		custPro.Addresses = addresses
	}()

	if ok := <-c; !ok {
		return nil, errors.ErrNotFoundUser
	}

	wg.Wait()
	return custPro, nil
}
