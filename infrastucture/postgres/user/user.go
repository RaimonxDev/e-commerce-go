package user

import (
	"context"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/jmoiron/sqlx"
)

var (
	psqlCreate     = `INSERT INTO users (id, email, password, details, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	psqlGetByEmail = `SELECT id, email , details, created_at, updated_at FROM users`
	psqlGetAll     = `SELECT id, email , details, created_at, updated_at FROM users where id = $1`
)

type User struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u User) Create(m *model.User) error {
	_, err := u.db.ExecContext(context.Background(), psqlCreate, m.ID, m.Email, m.Password, m.Details, m.CreatedAt, m.UpdatedAt)
	return err
}

func (u User) GetByEmail(email string) (model.User, error) {
	user := &model.User{}
	err := u.db.GetContext(context.Background(), user, psqlGetAll, email)
	return *user, err
}

func (u User) GetAll() (model.Users, error) {
	var users model.Users
	err := u.db.SelectContext(context.Background(), &users, psqlGetAll)
	return users, err
}
