package purchaseorder

import (
	"context"
	"fmt"

	"github.com/RaimonxDev/e-commerce-go.git/infrastructure/postgres"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const table = "purchase_orders"

var fields = []string{
	"id",
	"products",
	"created_at",
	"updated_at",
}

var (
	psqlCreate            = postgres.BuildSQLInsert(table, fields)
	psqlGetByID           = postgres.BuildSQLSelect(table, fields) + " WHERE id = $1"
	psqlValidatesProducts = postgres.BuildSQLSelect(table, fields) + " WHERE id = $1"
)

type PurchaseOrder struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *PurchaseOrder {
	return &PurchaseOrder{db: db}
}

func (po *PurchaseOrder) Create(purchaseOrder *model.PurchaseOrder) error {

	_, err := po.db.ExecContext(context.Background(), psqlCreate,
		purchaseOrder.ID,
		purchaseOrder.Products,
		purchaseOrder.Created_at,
		postgres.Int64ToNull(purchaseOrder.Updated_at)) // Convert int64 to null.Int64

	if err != nil {
		return fmt.Errorf("Error creating purchase order in database : %w", err)
	}

	return nil
}

func (po *PurchaseOrder) GetByID(id uuid.UUID) (model.PurchaseOrder, error) {
	purchaseOrder := model.PurchaseOrder{}
	err := po.db.GetContext(context.Background(), &purchaseOrder, psqlGetByID, id)
	return purchaseOrder, err

}

func (po *PurchaseOrder) ValidatesProducts(id uuid.UUID) (model.ValidatedProduct, error) {
	product := model.ValidatedProduct{}
	err := po.db.GetContext(context.Background(), &product, psqlValidatesProducts, id)
	return product, err
}
