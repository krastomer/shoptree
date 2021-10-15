package models

import "time"

type Product struct {
	ID             uint32
	Name           string
	ScientificName string
	Price          float32
	Description    string
	Status         ProductStatusType
	CreatedAt      time.Time
}

type ProductStatusType string

const (
	Unavailable ProductStatusType = "Unavailable"
	Available   ProductStatusType = "Available"
	Pending     ProductStatusType = "Pending"
	Purchased   ProductStatusType = "Purchased"
)

type ProductResponse struct {
	ID             uint32            `json:"id"`
	Name           string            `json:"name"`
	ScientificName string            `json:"scientific_name"`
	Price          float32           `json:"price"`
	Description    string            `json:"description"`
	Status         ProductStatusType `json:"status"`
}
