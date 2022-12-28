package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Price       int             `json:"price"`
	Images      json.RawMessage `json:"images"`
	Description string          `json:"description"`
	Features    json.RawMessage `json:"features"`
	CreatedAt   int64           `json:"created_at"`
	UpdatedAt   int64           `json:"updated_at"`
}

// Products is a slice of Product
type Products []Product
