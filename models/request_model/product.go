package request_model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SellerID    uuid.UUID `json:"seller_id"`
	CategoryID  uuid.UUID `json:"category_id"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductAddParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SellerID    uuid.UUID `json:"seller_id"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type ProductEditParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type ProductDeleteParams struct {
	ID        uuid.UUID `json:"id"`
	IsDeleted bool      `json:"is_deleted"`
}

