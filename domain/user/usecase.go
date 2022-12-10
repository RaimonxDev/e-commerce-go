package user

import (
	"fmt"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Repository Repository
}

func New(r Repository) User {
	return User{Repository: r}
}

func (u *User) Create(m model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error creating uuid: %w", err)
	}
	// Add ID to model
	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error bcrytp generate password: %w", err)
	}
	m.Password = string(password)

	// If details is nil o empty, create of empty json
	if m.Details == nil {
		m.Details = []byte("{}")
	}

	// TIMESTAMP
	m.CreatedAt = time.Now().Unix()
	m.UpdatedAt = time.Now().Unix()

	err = u.Repository.Create(&m)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (u *User) GetByEmail(email string) (*model.User, error) {
	return u.Repository.GetByEmail(email)
}

func (u *User) GetAll() (model.Users, error) {
	return u.Repository.GetAll()
}
