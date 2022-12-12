package user

import (
	"github.com/RaimonxDev/e-commerce-go.git/domain/user"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	userCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{userCase: uc}
}

func (h *handler) Create(c echo.Context) error {
	m := model.User{}
	// Bind request to model
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// Create user
	if err := h.userCase.Create(&m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error creating user"})
	}
	return c.JSON(http.StatusCreated, m)

}

func (h *handler) GetAll(c echo.Context) error {
	users, err := h.userCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error getting users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *handler) GetByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := h.userCase.GetByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error "})
	}
	return c.JSON(http.StatusOK, user)
}
