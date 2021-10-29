package auth

import "time"

type User struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Level       string `json:"level"`
}

type Customer struct {
	ID          uint32
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	BagLevel    uint8
	CreatedAt   time.Time
}

type Employee struct {
	ID          uint32
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Level       EmployeeLevelType
	CreatedAt   string
}

type EmployeeLevelType string

const (
	Admin   EmployeeLevelType = "Admin"
	Staff   EmployeeLevelType = "Staff"
	Deliver EmployeeLevelType = "Deliver"
)

type AuthRepository interface {
	GetCustomerByEmail(string) (*Customer, error)
	GetCustomerByPhone(string) (*Customer, error)
	GetEmployeeByEmail(string) (*Employee, error)
	RegisterCustomer(*Customer) error
}

type AuthService interface {
	Login(string, string) (string, error)
	Register(*User) error
}
