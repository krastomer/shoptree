package search

import (
	"context"
	"database/sql"
	"strings"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsCategoryProductMR = &dbq.Options{ConcreteStruct: CategoryProduct{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductMR         = &dbq.Options{ConcreteStruct: Product{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_CATEGORIES_PRODUCT = "SELECT * FROM `categories`;"
	QUERY_GET_PRODUCTS_LIKE      = "SELECT * FROM products WHERE name LIKE '%?%';"
)

func NewSearchRepository(db *sql.DB) SearchRepository {
	return &repository{db: db}
}

func (r *repository) GetCategoriesProduct(ctx context.Context) (cat []*CategoryProduct, err error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_CATEGORIES_PRODUCT, OptsCategoryProductMR)
	cat = result.([]*CategoryProduct)
	if len(cat) == 0 {
		return nil, sql.ErrNoRows
	}

	return cat, nil
}

func (r *repository) GetProductsLike(ctx context.Context, data string) (products []*Product, err error) {
	query := strings.Replace(QUERY_GET_PRODUCTS_LIKE, "?", data, 1)

	result := dbq.MustQ(ctx, r.db, query, OptsProductMR)
	products = result.([]*Product)
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}
	return products, nil
}
