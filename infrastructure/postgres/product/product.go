package product

import (
	"context"
	"fmt"
	"github.com/RaimonxDev/e-commerce-go.git/infrastructure/postgres"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const table = "products"

var fields = []string{
	"id",
	"name",
	"description",
	"price",
	"stock",
	"images",
	"features",
	"created_at",
	"updated_at",
}

var (
	psqlCreateProduct = postgres.BuildSQLInsert(table, fields)
	psqlUpdateProduct = postgres.BuildSQLUpdateByID(table, fields)
	psqlDeleteProduct = postgres.BuildSQLDelete(table)
	psqlGetByID       = postgres.BuildSQLSelect(table, fields) + " WHERE id = $1"
	psqlGetAll        = postgres.BuildSQLSelect(table, fields)
)

type Product struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Product {
	return &Product{db: db}
}

func (p Product) Create(m *model.Product) error {
	_, err := p.db.ExecContext(context.Background(), psqlCreateProduct,
		m.ID,
		m.Name,
		m.Description,
		m.Price,
		m.Images,
		m.Features,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	return err
}

func (p Product) Update(m *model.Product) error {
	_, err := p.db.ExecContext(context.Background(), psqlUpdateProduct,
		m.ID,
		m.Name,
		m.Description,
		m.Price,
		m.Images,
		m.Features,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)

	if err != nil {
		fmt.Errorf("error updating product in db: %v", err)
	}

	return nil
}

func (p Product) Delete(id uuid.UUID) error {

	_, err := p.db.ExecContext(context.Background(), psqlDeleteProduct, id)
	return err
}

func (p Product) GetAll() (model.Products, error) {
	var pp model.Products
	err := p.db.SelectContext(context.Background(), &pp, psqlGetAll)
	return pp, err
}

func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {

	product := model.Product{}

	err := p.db.GetContext(context.Background(), &product, psqlGetByID, ID)

	return product, err

}
