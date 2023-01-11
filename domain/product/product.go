package product

import (
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
)

type UseCase interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetAll(pagination model.Pagination) (model.Products, error)
	GetByID(ID uuid.UUID) (model.Product, error)
}

type Repository interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetAll(pagination model.Pagination) (model.Products, error)
	GetByID(ID uuid.UUID) (model.Product, error)
}
