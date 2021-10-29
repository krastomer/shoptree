package mariadb

import (
	"github.com/krastomer/shoptree/backend/pkg/product"
	"gorm.io/gorm"
)

const (
	QUERY_GET_PRODUCT_BY_ID = "SELECT * FROM `products` WHERE id = ?"
	QUERY_ADD_PRODUCT       = "INSERT INTO `products` (`name`, `scientific_name`, `price`, `description`, `status`) VALUES (?, ?, ?, ?, ?);"
)

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetProductByID(id uint32) (*product.Product, error) {
	prod := &product.Product{}
	row := r.db.Raw(QUERY_GET_PRODUCT_BY_ID, id).Row()
	row.Scan(
		&prod.ID,
		&prod.Name,
		&prod.ScientificName,
		&prod.Price,
		&prod.Description,
		&prod.Status,
		&prod.CreatedAt,
	)
	if prod.ID == 0 {
		return nil, ErrQueryNotFound
	}
	return prod, nil
}

func (r *mariaDBRepository) AddProduct(product *product.Product) error {
	result := r.db.Exec(
		QUERY_ADD_PRODUCT,
		product.Name,
		product.ScientificName,
		product.Price,
		product.Description,
		product.Status,
	)
	if result.Error != nil {
		return ErrInsertFailed
	}
	return nil
}
