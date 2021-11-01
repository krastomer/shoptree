package product

import (
	"errors"

	"gorm.io/gorm"
)

type mariaDBRepository struct {
	db *gorm.DB
}

const (
	QUERY_GET_PRODUCT_BY_ID     = "SELECT * FROM `products` WHERE id = ?"
	QUERY_GET_PRODUCT_IMAGES_ID = "SELECT `product_images`.id FROM `products` JOIN `product_images` ON `products`.id = `product_images`.product_id WHERE `products`.id = ?;"
)

var (
	ErrQueryNotFound       = errors.New("query not found")
	ErrInternalServerError = errors.New("internal server error")
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetProductByID(id int) (*Product, error) {
	prod := &Product{}
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
	// TODO: change to check with Time
	if prod.ID == 0 {
		return nil, ErrQueryNotFound
	}
	return prod, nil
}

func (r *mariaDBRepository) GetProductImagesID(id int) ([]int, error) {
	var imagesID []int
	rows, err := r.db.Raw(QUERY_GET_PRODUCT_IMAGES_ID, id).Rows()
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var imageID int
		rows.Scan(&imageID)
		imagesID = append(imagesID, imageID)
	}

	return imagesID, nil
}
