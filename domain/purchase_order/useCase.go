package purchaseorder

import (
	"encoding/json"
	"fmt"

	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
)

type PurchaseOrder struct {
	repo Repository
}

// NewPurchaseOrder returns a new instance of the purchase order
func NewPurchaseOrder(repo Repository) *PurchaseOrder {
	return &PurchaseOrder{repo: repo}
}

func (p *PurchaseOrder) Create(m *model.PurchaseOrder) error {
	// Validate si los datos del modelo producto son correctos y no estan vacios
	if err := m.Validate(); err != nil {
		return fmt.Errorf("purchase order is not valid: %w", err)
	}

	selectedProducts := []model.ProductToPurchase{}
	if err := json.Unmarshal(m.Products, &selectedProducts); err != nil {
		return fmt.Errorf("json.Unmarshal() in purchase order: %w", err)
	}
	// Validate products in the purchase order and return the validated products
	validatedProducts, err := p.ValidatesProducts(selectedProducts)
	if err != nil {
		return fmt.Errorf("error validating products: %w", err)
	}

	for _, v := range validatedProducts {
		for _, sp := range selectedProducts {
			if v.ProductID == sp.ProductID {
				sp.UnitPrice = v.UnitPrice
			}
		}
	}

	pp, err := json.Marshal(selectedProducts)
	if err != nil {
		return fmt.Errorf("json.Marshal() in purchase order: %w", err)
	}
	// Set the validated products in the purchase order
	m.Products = pp
	ID, err := uuid.NewUUID()

	if err != nil {
		return fmt.Errorf("error generating uuid: %w", err)
	}
	// Set the ID of the purchase order
	m.ID = ID

	if err := p.repo.Create(m); err != nil {
		return fmt.Errorf("error creating purchase order: %w", err)
	}

	return nil
}

func (p *PurchaseOrder) GetByID(id uuid.UUID) (model.PurchaseOrder, error) {
	return p.repo.GetByID(id)
}

func (p *PurchaseOrder) ValidatesProducts(pp []model.ProductToPurchase) ([]model.ValidatedProduct, error) {
	data := []model.ValidatedProduct{}
	for _, v := range pp {
		p, err := p.repo.ValidatesProducts(v.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product not found: %w", err)
		}
		data = append(data, p)
	}
	return data, nil
}
