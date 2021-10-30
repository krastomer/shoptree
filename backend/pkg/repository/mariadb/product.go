package mariadb

import (
	"github.com/krastomer/shoptree/backend/pkg/product"
	"gorm.io/gorm"
)

const (
	QUERY_GET_PRODUCT_BY_ID       = "SELECT * FROM `products` WHERE id = ?"
	QUERY_GET_PRODUCTS            = "SELECT * FROM `products` ORDER BY status LIMIT 20"
	QUERY_GET_PRODUCT_IMAGE_BY_ID = "SELECT `image_path` FROM `product_images` WHERE id = ?"
	QUERY_GET_PRODUCT_IMAGES_ID   = "SELECT `product_images`.id FROM `products` JOIN `product_images` ON `products`.id = `product_images`.product_id WHERE `products`.id = ?;"
	QUERY_ADD_PRODUCT             = "INSERT INTO `products` (`name`, `scientific_name`, `price`, `description`, `status`) VALUES (?, ?, ?, ?, ?);"
	QUERY_ADD_PRODUCT_IMAGE       = "INSERT INTO `product_images` (`product_id`, `image_path`) VALUES (?, ?);"
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

func (r *mariaDBRepository) GetProducts() ([]*product.Product, error) {
	var products []*product.Product
	rows, err := r.db.Raw(QUERY_GET_PRODUCTS).Rows()
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		product := &product.Product{}
		r.db.ScanRows(rows, product)
		products = append(products, product)
	}
	return products, nil
}

func (r *mariaDBRepository) AddProductImage(id uint32, path string) error {
	result := r.db.Exec(
		QUERY_ADD_PRODUCT_IMAGE,
		id,
		path,
	)
	if result.Error != nil {
		return ErrInsertFailed
	}
	return nil
}

func (r *mariaDBRepository) GetProductImageByID(id uint32) (string, error) {
	var result string
	row := r.db.Raw(QUERY_GET_PRODUCT_IMAGE_BY_ID, id).Row()
	row.Scan(&result)
	if result == "" {
		return "", ErrQueryNotFound
	}
	return result, nil
}

func (r *mariaDBRepository) GetProductImagesID(id uint32) ([]uint32, error) {
	var imagesID []uint32
	rows, err := r.db.Raw(QUERY_GET_PRODUCT_IMAGES_ID, id).Rows()
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var imageID uint32
		rows.Scan(&imageID)
		imagesID = append(imagesID, imageID)
	}

	return imagesID, nil
}
