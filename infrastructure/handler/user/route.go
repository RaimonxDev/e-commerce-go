package user

import (
	"github.com/RaimonxDev/e-commerce-go.git/domain/user"
	repositoryUser "github.com/RaimonxDev/e-commerce-go.git/infrastructure/postgres/user"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, dbPool *sqlx.DB) {
	h := buildHandler(dbPool)
	adminRoutes(e, h)
	publicRoutes(e, h)
}

func buildHandler(dbPool *sqlx.DB) handler {
	repository := repositoryUser.New(dbPool)
	useCase := user.New(repository)
	return newHandler(&useCase)
}

func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/users")
	g.GET("", h.GetAll)

}
func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")
	g.POST("", h.Create)
}
