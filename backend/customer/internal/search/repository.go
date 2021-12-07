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
	OptsProductSR         = &dbq.Options{ConcreteStruct: Product{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsCategoryProductMR = &dbq.Options{ConcreteStruct: CategoryProduct{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductMR         = &dbq.Options{ConcreteStruct: Product{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductPendingSR  = &dbq.Options{ConcreteStruct: ProductPending{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_CATEGORIES_PRODUCT      = "SELECT * FROM `categories`;"
	QUERY_GET_PRODUCTS_LIKE           = "SELECT * FROM products WHERE name LIKE '%?%';"
	QUERY_GET_IMAGE_PRODUCT_BY_ID     = "SELECT id FROM `images_product` WHERE product_id = ? LIMIT 1;"
	QUERY_GET_PRODUCT_AVAILABLE_BY_ID = "SELECT * FROM `products_available` WHERE id = ?;"
	QUERY_GET_PRODUCT_PENDING_BY_ID   = "SELECT * FROM `products_pending` WHERE product_id = ?;"
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
func (r *repository) GetImageProductByID(ctx context.Context, id int) (path int, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_IMAGE_PRODUCT_BY_ID, dbq.SingleResult, args)
	if result == nil {
		return -1, sql.ErrNoRows
	}
	path = int(result.(map[string]interface{})["id"].(int32))
	return path, nil
}

func (r *repository) GetProductAvailableByID(ctx context.Context, id int) (prod *Product, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_AVAILABLE_BY_ID, OptsProductSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	prod = result.(*Product)
	return prod, nil
}

func (r *repository) GetProductPendingByID(ctx context.Context, id int) (prod *ProductPending, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_PENDING_BY_ID, OptsProductPendingSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	prod = result.(*ProductPending)
	return prod, nil
}
