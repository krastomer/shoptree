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
	OptsProductSR         = &dbq.Options{ConcreteStruct: Product{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductMR         = &dbq.Options{ConcreteStruct: Product{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsCategoryProductMR = &dbq.Options{ConcreteStruct: CategoryProduct{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsImageProductMR    = &dbq.Options{ConcreteStruct: ImageProduct{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductPendingSR  = &dbq.Options{ConcreteStruct: ProductPending{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_PRODUCTS                = "SELECT * FROM `products`;"
	QUERY_GET_PRODUCT_BY_ID           = "SELECT * FROM `products` WHERE id = ?;"
	QUERY_GET_CATEGORIES_PRODUCT      = "SELECT * FROM `categories_product_name` WHERE product_id = ?;"
	QUERY_GET_IMAGES_PRODUCT_ID       = "SELECT * FROM `images_product` WHERE product_id = ?;"
	QUERY_GET_PRODUCT_AVAILABLE_BY_ID = "SELECT * FROM `products_available` WHERE id = ?;"
	QUERY_GET_PRODUCT_PENDING_BY_ID   = "SELECT * FROM `products_pending` WHERE product_id = ?;"
	QUERY_GET_IMAGE_PRODUCT_BY_ID     = "SELECT `image_path` FROM `images_product` WHERE id = ?"
)

func NewProductRepository(db *sql.DB) ProductRepository {
	return &repository{db: db}
}

func (r *repository) GetProducts(ctx context.Context) (prod []*Product, err error) {
	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCTS, OptsProductMR)
	prod = result.([]*Product)
	if len(prod) == 0 {
		return nil, sql.ErrNoRows
	}
	return prod, nil
}

func (r *repository) GetProductByID(ctx context.Context, id int) (prod *Product, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_BY_ID, OptsProductSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	prod = result.(*Product)
	return prod, nil
}

func (r *repository) GetCategoriesProduct(ctx context.Context, id int) (cat []*CategoryProduct, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CATEGORIES_PRODUCT, OptsCategoryProductMR, args)
	cat = result.([]*CategoryProduct)
	if len(cat) == 0 {
		return nil, sql.ErrNoRows
	}

	return cat, nil
}

func (r *repository) GetImagesProductID(ctx context.Context, id int) (images []*ImageProduct, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_IMAGES_PRODUCT_ID, OptsImageProductMR, args)
	images = result.([]*ImageProduct)
	if len(images) == 0 {
		return nil, sql.ErrNoRows
	}

	return images, nil
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

func (r *repository) GetImageProductByID(ctx context.Context, id int) (path string, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_IMAGE_PRODUCT_BY_ID, dbq.SingleResult, args)
	if p, ok := result.(map[string]interface{}); ok {
		path = p["image_path"].(string)
		return path, nil
	}
	return "", sql.ErrNoRows
}
