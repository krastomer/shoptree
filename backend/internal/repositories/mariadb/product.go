package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_PRODUCT = "SELECT * FROM `products` WHERE id = ?"
)

func NewProductRepo(db *gorm.DB) entities.ProductRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetProductByID(id int) (*models.Product, error) {
	product := &models.Product{}
	row := r.db.Raw(QUERY_GET_PRODUCT, id).Row()
	row.Scan(
		&product.ID,
		&product.Name,
		&product.ScienceName,
		&product.Price,
		&product.Description,
		&product.Status,
	)
	if product.Name == "" {
		return nil, errors.ErrQueryNotFound
	}

	return product, nil
}
