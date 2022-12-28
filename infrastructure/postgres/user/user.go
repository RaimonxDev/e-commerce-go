package user

import (
	"context"
	"github.com/RaimonxDev/e-commerce-go.git/infrastructure/postgres"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/jmoiron/sqlx"
)

const table = "users"

var fields = []string{
	"id",
	"email",
	"password",
	"details",
	"created_at",
	"updated_at",
}

var (
	psqlCreate     = postgres.BuildSQLInsert(table, fields)
	psqlGetByEmail = postgres.BuildSQLSelect(table, fields) + " WHERE email = $1"
	psqlGetAll     = postgres.BuildSQLSelect(table, fields)
)

type User struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u User) Create(m *model.User) error {
	_, err := u.db.ExecContext(context.Background(), psqlCreate,
		m.ID,
		m.Email,
		m.Password,
		m.Details,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt)) // Convert int64 to null.Int64
	return err
}

func (u User) GetByEmail(email string) (model.User, error) {
	user := &model.User{}
	err := u.db.GetContext(context.Background(), user, psqlGetByEmail, email)
	return *user, err
}

func (u User) GetAll() (model.Users, error) {
	var users model.Users
	err := u.db.SelectContext(context.Background(), &users, psqlGetAll)
	return users, err
}
