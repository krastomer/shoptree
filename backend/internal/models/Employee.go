package models

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
