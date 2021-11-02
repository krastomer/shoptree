package product

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type mariaDBRepository struct {
	db *gorm.DB
}

const (
	QUERY_GET_PRODUCT_BY_ID       = "SELECT * FROM `products` WHERE id = ?"
	QUERY_GET_PRODUCT_IMAGES_ID   = "SELECT `product_images`.id FROM `products` JOIN `product_images` ON `products`.id = `product_images`.product_id WHERE `products`.id = ?;"
	QUERY_GET_PRODUCT_IMAGE_BY_ID = "SELECT `image_path` FROM `product_images` WHERE id = ?"

	QUERY_CREATE_PRODUCT_IMAGE_PATH = "INSERT INTO `product_images` (`product_id`, `image_path`) VALUES (?, ?);"
	QUERY_CREATE_PRODUCT            = "INSERT INTO `products` (`name`, `scientific_name`, `price`, `description`, `status`) VALUES (?, ?, ?, ?, ?);"
)

var (
	ErrQueryNotFound       = errors.New("query not found")
	ErrInsertFailed        = errors.New("insert failed")
	ErrInternalServerError = errors.New("internal server error")
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &mariaDBRepository{db: db}
}

// TODO: fix
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

	fmt.Println(prod)
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

func (r *mariaDBRepository) GetProductImageByID(id int) (string, error) {
	var result string
	row := r.db.Raw(QUERY_GET_PRODUCT_IMAGE_BY_ID, id).Row()
	row.Scan(&result)
	if result == "" {
		return "", ErrQueryNotFound
	}
	return result, nil
}

func (r *mariaDBRepository) CreateProduct(product *ProductRequest) error {
	result := r.db.Exec(
		QUERY_CREATE_PRODUCT,
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

func (r *mariaDBRepository) CreateProductImagePath(image *ProductImageRequest) error {
	result := r.db.Exec(
		QUERY_CREATE_PRODUCT_IMAGE_PATH,
		image.ID,
		image.Path,
	)
	if result.Error != nil {
		return ErrInsertFailed
	}
	return nil
}
