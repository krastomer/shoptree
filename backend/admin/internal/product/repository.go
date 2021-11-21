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
	OptsProductSR = &dbq.Options{ConcreteStruct: Product{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_PRODUCTS         = "SELECT * FROM `products`;"
	QUERY_GET_PRODUCT_BY_ID    = "SELECT * FROM `products` WHERE id = ?;"
	QUERY_CREATE_PRODUCT       = "INSERT INTO `products` (`name`, `scientific_name`, `description`, `price`) VALUES (?, ?, ?, ?);"
	QUERY_UPDATE_PRODUCT       = "UPDATE `products` SET name = ?, scientific_name = ?, description = ?, price = ? WHERE id = ?;"
	QUERY_DELETE_PRODUCT_BY_ID = "DELETE FROM `products` WHERE id = ?;"
	QUERY_CREATE_IMAGE_PRODUCT = "INSERT INTO `images_product` (`product_id`, `image_path`) VALUES (?, ?);"
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

func (r *repository) GetProductByID(ctx context.Context, id int) (product *Product, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_BY_ID, OptsProductSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	product = result.(*Product)
	return product, nil
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

func (r *repository) UpdateProduct(ctx context.Context, product *Product) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_UPDATE_PRODUCT, nil,
			product.Name,
			product.ScientificName,
			product.Description,
			product.Price,
			product.ID,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}

func (r *repository) DeleteProductByID(ctx context.Context, id int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_DELETE_PRODUCT_BY_ID, nil,
			id,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}

func (r *repository) CreateImageProduct(ctx context.Context, image *ImageProduct) (err error) {
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
