package auth

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsEmployeeSR = &dbq.Options{ConcreteStruct: Employee{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?;"
)

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &repository{db: db}
}

func (r *repository) GetEmployeeByEmail(ctx context.Context, email string) (empl *Employee, _ error) {
	args := []interface{}{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_EMPLOYEE_BY_EMAIL, OptsEmployeeSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	empl = result.(*Employee)
	return empl, nil
}
