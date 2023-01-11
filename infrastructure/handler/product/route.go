package product

import (
	"github.com/RaimonxDev/e-commerce-go.git/domain/product"
	repo "github.com/RaimonxDev/e-commerce-go.git/infrastructure/postgres/product"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, db *sqlx.DB) {
	h := buildHandler(db)
	adminRoutes(e, h)
	publicRoutes(e, h)
}

func buildHandler(db *sqlx.DB) handler {
	repo := repo.NewRepository(db)
	useCase := product.NewProductUseCase(repo)
	return newHandler(&useCase)
}
func adminRoutes(e *echo.Echo, h handler) {

	route := e.Group("/api/v1/admin/products")
	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/products")
	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}
