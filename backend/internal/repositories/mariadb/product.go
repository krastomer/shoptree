package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_PRODUCT_BY_ID = "SELECT * FROM `products` WHERE id = ?"
)

func NewProductRepo(db *gorm.DB) entities.ProductRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetProduct(id uint32) (*models.Product, error) {
	prod := &models.Product{}
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
		return nil, errors.ErrQueryNotFound
	}
	return prod, nil
}
