package user

import "github.com/RaimonxDev/e-commerce-go.git/model"

// Datos de entrada para el caso de uso
type UseCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

// Datos de salidad para el caso de uso
type Repository interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
