package profile

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
)

type profileService struct {
	repo ProfileRepository
}

func NewProfileService(repo ProfileRepository) ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) GetProfileCustomer(id uint32) (*CustomerProfile, error) {
	custPro := &CustomerProfile{}

	cust, err := s.repo.GetCustomerByID(id)
	if err != nil {
		return nil, ErrCustomerNotFound
	}

	addresses, _ := s.repo.GetAddresses(id)

	{
		custPro.Name = cust.Name
		custPro.Email = cust.Email
		custPro.PhoneNumber = cust.PhoneNumber
		custPro.CreatedAt = cust.CreatedAt
		custPro.Addresses = addresses
	}

	return custPro, nil
}
