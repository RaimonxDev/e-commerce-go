package model

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/google/uuid"
)

type PurchaseOrder struct {
	ID         uuid.UUID       `json:"id"`
	UserID     uuid.UUID       `json:"user_id"`
	Products   json.RawMessage `json:"products"`
	Created_at int64           `json:"created_at"`
	Updated_at int64           `json:"updated_at"`
}

func (p PurchaseOrder) Validate() error {
	if p.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if len(p.Products) == 0 {
		return errors.New("products is empty")
	}
	selectedProducts := []ProductToPurchase{}

	if err := json.Unmarshal(p.Products, &selectedProducts); err != nil {
		log.Fatalf("%s %v", "json.Unmarshal()", err)
	}

	for _, v := range selectedProducts {
		if v.ProductID == uuid.Nil {
			return errors.New("product_id is required")
		}
		if v.Amount < 1 {
			return errors.New("the amount must be greater than 1")
		}
		if v.UnitPrice < 0 {
			return errors.New("the unit price must be greater than 0")
		}
	}
	return nil

}

type ProductToPurchase struct {
	ProductID uuid.UUID `json:"product_id"`
	UnitPrice float64   `json:"unit_price"`
	Amount    float64   `json:"amount"`
}
