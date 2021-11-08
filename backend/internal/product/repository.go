package product

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type mariaDBRepository struct {
	db *sql.DB
}

var (
	OptsProductSR         = &dbq.Options{ConcreteStruct: Product{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsCategoryProductMR = &dbq.Options{ConcreteStruct: CategoryProduct{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsImageProductMR    = &dbq.Options{ConcreteStruct: ImageProduct{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_PRODUCT_BY_ID       = "SELECT * FROM `products` WHERE id = ?"
	QUERY_GET_CATEGORIES_PRODUCT  = "SELECT * FROM `categories_product_name` WHERE product_id = ?"
	QUERY_GET_IMAGE_PRODUCT_BY_ID = "SELECT `image_path` FROM `images_product` WHERE id = ?"
	QUERY_GET_IMAGES_PRODUCT_ID   = "SELECT * FROM `images_product` WHERE product_id = ?"

	QUERY_CREATE_IMAGE_PRODUCT = "INSERT INTO `images_product` (`product_id`, `image_path`) VALUES (?, ?);"
	QUERY_CREATE_PRODUCT       = "INSERT INTO `products` (`name`, `scientific_name`, `price`, `description`, `status`) VALUES (?, ?, ?, ?, ?);"
)

func NewProductRepository(db *sql.DB) ProductRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetProductByID(ctx context.Context, id int) (prod *Product, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_BY_ID, OptsProductSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	prod = result.(*Product)
	return prod, nil
}

func (r *mariaDBRepository) GetCategoriesProduct(ctx context.Context, id int) (cat []*CategoryProduct, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CATEGORIES_PRODUCT, OptsCategoryProductMR, args)
	cat = result.([]*CategoryProduct)
	if len(cat) == 0 {
		return nil, sql.ErrNoRows
	}

	return cat, nil
}

func (r *mariaDBRepository) GetImagesProductID(ctx context.Context, id int) (images []*ImageProduct, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_IMAGES_PRODUCT_ID, OptsImageProductMR, args)
	images = result.([]*ImageProduct)
	if len(images) == 0 {
		return nil, sql.ErrNoRows
	}

	return images, nil
}

func (r *mariaDBRepository) GetImageProductByID(ctx context.Context, id int) (path string, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_IMAGE_PRODUCT_BY_ID, dbq.SingleResult, args)
	if p, ok := result.(map[string]interface{}); ok {
		path = p["image_path"].(string)
		return path, nil
	}
	return "", sql.ErrNoRows
}

// func (r *mariaDBRepository) CreateProduct(product *ProductRequest) error {
// 	result := r.db.Exec(
// 		QUERY_CREATE_PRODUCT,
// 		product.Name,
// 		product.ScientificName,
// 		product.Price,
// 		product.Description,
// 		product.Status,
// 	)
// 	if result.Error != nil {
// 		return ErrInsertFailed
// 	}
// 	return nil
// }

func (r *mariaDBRepository) CreateImageProduct(ctx context.Context, image *ImageProduct) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_CREATE_IMAGE_PRODUCT, nil,
			image.ProductID,
			image.ImagePath,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}
