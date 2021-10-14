package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?"
)

func NewEmployeeRepo(db *gorm.DB) entities.EmployeeRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetEmployeeByEmail(email string) (*models.Employee, error) {
	empl := &models.Employee{}
	row := r.db.Raw(QUERY_GET_EMPLOYEE_BY_EMAIL, email).Row()
	row.Scan(
		&empl.ID,
		&empl.Name,
		&empl.Email,
		&empl.Password,
		&empl.PhoneNumber,
		&empl.Level,
		&empl.CreatedAt,
	)
	if empl.Name == "" {
		return nil, errors.ErrQueryNotFound
	}
	return empl, nil
}
