package customer

import (
	"context"
	"database/sql"
)

type mariaDBRepository struct {
	db *sql.DB
}

const (
	QUERY_CREATE_CUSTOMER        = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`) VALUES (?, ?, ?, ?);"
	QUERY_GET_CUSTOMER_BY_EMAIL  = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_CUSTOMER_BY_PHONE  = "SELECT * FROM `customers` WHERE phone_number = ?"
	QUERY_GET_ADDRESSES_CUSTOMER = "SELECT * FROM `addresses_customer` WHERE customer_id = ?"

	QUERY_CREATE_ADDRESS     = "INSERT INTO `address_customers` ( `customer_id`, `name`, `phone_number`, `address_line`, `country`, `state`, `city`, `district`, `postal_code`) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	QUERY_GET_CUSTOMER_BY_ID = "SELECT * FROM `customers` WHERE id = ?"

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

func (r *mariaDBRepository) GetCustomerByID(ctx context.Context, id int) (*Customer, error) {
	cust := &Customer{}

	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_CUSTOMER_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, id).Scan(
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

func (r *mariaDBRepository) GetAddressesCustomer(ctx context.Context, id int) ([]*Address, error) {
	var addresses []*Address

	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_ADDRESSES_CUSTOMER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := r.db.QueryContext(ctx, QUERY_GET_ADDRESSES_CUSTOMER, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		address := &Address{}
		err = res.Scan(
			&address.ID,
			&address.CustomerID,
			&address.Name,
			&address.PhoneNumber,
			&address.AddressLine,
			&address.Country,
			&address.State,
			&address.City,
			&address.District,
			&address.PostalCode,
			&address.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
