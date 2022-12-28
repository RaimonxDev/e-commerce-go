package product

import (
	"fmt"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Repository Repository
}

func new(r Repository) Product {
	return Product{Repository: r}
}

func (p *Product) Create(m *model.Product) error {

	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error creating uuid: %w", err)
	}
	// Add ID to model
	m.ID = ID

	// If details is nil o empty, create of empty array
	if len(m.Images) == 0 {
		m.Images = []byte("[]")
	}
	if len(m.Features) == 0 {
		m.Features = []byte("{}")
	}
	// TIMESTAMP
	m.CreatedAt = time.Now().Unix()

	err = p.Repository.Create(m)
	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	return nil

}

func (p *Product) Update(m *model.Product) error {

	if m.ID == uuid.Nil {
		return fmt.Errorf("error product id no valid or empty")
	}
	if len(m.Images) == 0 {
		m.Images = []byte("[]")
	}
	if len(m.Features) == 0 {
		m.Features = []byte("{}")
	}

	m.UpdatedAt = time.Now().Unix()
	err := p.Repository.Update(m)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil

}

func (p *Product) Delete(ID uuid.UUID) error {
	if ID == uuid.Nil {
		return fmt.Errorf("error product id no valid or empty")
	}
	err := p.Repository.Delete(ID)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}

func (p *Product) GetAll() (model.Products, error) {
	return p.Repository.GetAll()
}

func (p *Product) GetByID(ID uuid.UUID) (model.Product, error) {

	if ID == uuid.Nil {
		fmt.Errorf("error product id no valid or empty")
	}

	return p.Repository.GetByID(ID)
}
