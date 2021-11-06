package customer

import (
	"context"
	"database/sql"
)

type mariaDBRepository struct {
	db *sql.DB
}

const (
	QUERY_CREATE_CUSTOMER       = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`) VALUES (?, ?, ?, ?);"
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_CUSTOMER_BY_PHONE = "SELECT * FROM `customers` WHERE phone_number = ?"

	QUERY_CREATE_ADDRESS     = "INSERT INTO `address_customers` ( `customer_id`, `name`, `phone_number`, `address_line`, `country`, `state`, `city`, `district`, `postal_code`) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	QUERY_GET_CUSTOMER_BY_ID = "SELECT * FROM `customers` WHERE id = ?"
	QUERY_GET_ADDRESSES      = "SELECT * FROM `address_customers` WHERE customer_id = ?"

	QUERY_GET_INVOICES = "SELECT * FROM `invoices` WHERE customer_id = ?;"
)

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) CreateCustomer(ctx context.Context, cust *CustomerRequest) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, QUERY_CREATE_CUSTOMER, cust.Name, cust.Email, cust.Password, cust.PhoneNumber)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *mariaDBRepository) GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	cust := &Customer{}

	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_CUSTOMER_BY_EMAIL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, email).Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return cust, nil
}

func (r *mariaDBRepository) GetCustomerByPhone(ctx context.Context, phone string) (*Customer, error) {
	cust := &Customer{}

	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_CUSTOMER_BY_PHONE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, phone).Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return cust, nil
}

// func (r *mariaDBRepository) GetAddresses(id int) ([]*Address, error) {
// 	var addresses []*Address
// 	rows, err := r.db.Raw(QUERY_GET_ADDRESSES, id).Rows()
// 	if err != nil {
// 		return nil, ErrInternalServerError
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		address := &Address{}
// 		r.db.ScanRows(rows, address)
// 		addresses = append(addresses, address)
// 	}
// 	return addresses, nil
// }

// func (r *mariaDBRepository) GetCustomerByID(id int) (*Customer, error) {
// 	cust := &Customer{}
// 	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_ID, id).Row()
// 	row.Scan(
// 		&cust.ID,
// 		&cust.Name,
// 		&cust.Email,
// 		&cust.Password,
// 		&cust.PhoneNumber,
// 		&cust.CreatedAt,
// 	)

// 	// TODO : change to compare with time
// 	if cust.Name == "" {
// 		return nil, ErrQueryNotFound
// 	}
// 	return cust, nil
// }

// func (r *mariaDBRepository) CreateAddress(id int, address *Address) error {
// 	result := r.db.Exec(
// 		QUERY_CREATE_ADDRESS,
// 		id,
// 		address.Name,
// 		address.PhoneNumber,
// 		address.AddressLine,
// 		address.Country,
// 		address.State,
// 		address.City,
// 		address.District,
// 		address.PostalCode,
// 	)

// 	if result.Error != nil {
// 		return ErrInsertFailed
// 	}
// 	return nil
// }

// func (r *mariaDBRepository) GetInvoices(id int) ([]*Order, error) {
// 	var orders []*Order
// 	rows, err := r.db.Raw(QUERY_GET_INVOICES, id).Rows()
// 	if err != nil {
// 		return nil, ErrInternalServerError
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		order := &Order{}
// 		r.db.ScanRows(rows, order)
// 		orders = append(orders, order)
// 	}
// 	return orders, nil
// }
