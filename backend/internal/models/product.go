package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ScienceName string  `json:"science_name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}
