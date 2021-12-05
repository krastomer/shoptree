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
	QUERY_GET_CATEGORIES_PRODUCT       = "SELECT * FROM `categories`;"
	QUERY_GET_PRODUCTS_LIKE            = "SELECT * FROM products WHERE name LIKE \"%%?%%\";"
	QUERY_GET_IMAGE_PRODUCT_BY_ID      = "SELECT id FROM `images_product` WHERE product_id = ? LIMIT 1;"
	QUERY_CATEGORIES_PRODUCT_ID_BY_CAT = "SELECT categories_product.product_id FROM categories_product JOIN categories ON categories.id = categories_product.category_id WHERE categories.name = ?"
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
	query := strings.ReplaceAll(QUERY_GET_PRODUCTS_LIKE, "?", data)

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

func (r *repository) GetCategoriesProducts(ctx context.Context, cat string) (list []int, err error) {
	args := []interface{}{cat}

	result := dbq.MustQ(ctx, r.db, QUERY_CATEGORIES_PRODUCT_ID_BY_CAT, nil, args)

	data := result.([]map[string]interface{})
	if len(data) == 0 {
		return nil, sql.ErrNoRows
	}
	for _, d := range data {
		list = append(list, int(d["product_id"].(int32)))
	}

	return list, nil
}
