package purchaseorder

import (
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
)

type UsesCase interface {
	Create(purchaseOrder *model.PurchaseOrder) error
	GetByID(id uuid.UUID) (model.PurchaseOrder, error)
}

type Repository interface {
	Create(purchaseOrder *model.PurchaseOrder) error
	GetByID(id uuid.UUID) (model.PurchaseOrder, error)
	ValidatesProducts(productID uuid.UUID) (model.ValidatedProduct, error)
}
