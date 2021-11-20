package product

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsProductMR = &dbq.Options{ConcreteStruct: Product{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_PRODUCTS   = "SELECT * FROM `products`;"
	QUERY_CREATE_PRODUCT = "INSERT INTO `products` (`name`, `scientific_name`, `description`, `price`) VALUES (?, ?, ?, ?);"
)

func NewProductRepository(db *sql.DB) ProductRepository {
	return &repository{db: db}
}

func (r *repository) GetProducts(ctx context.Context) (products []*Product, _ error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCTS, OptsProductMR)

	products = result.([]*Product)
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}

	return products, nil
}

func (r *repository) CreateProduct(ctx context.Context, product *Product) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_CREATE_PRODUCT, nil,
			product.Name,
			product.ScientificName,
			product.Description,
			product.Price,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}
