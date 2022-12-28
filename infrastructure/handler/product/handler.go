package product

import (
	"github.com/RaimonxDev/e-commerce-go.git/domain/product"
	"github.com/RaimonxDev/e-commerce-go.git/infrastructure/handler/response"
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"strconv"
)

type handler struct {
	useCase  product.UseCase
	response response.API
}

func newHandler(useCase product.UseCase) handler {
	return handler{useCase: useCase}
}

func (h *handler) Create(c echo.Context) error {

	p := model.Product{}
	// Bind request to model
	err := c.Bind(&p)
	if err != nil {
		return h.response.BindFailed(err)
	}
	// Create product

	err = h.useCase.Create(&p)

	if err != nil {
		return h.response.Error(c, "UseCase Create()", err)
	}
	// devolver respuesta al usuario
	return c.JSON(h.response.Created(p))

}

func (h *handler) Update(c echo.Context) error {
	p := model.Product{}
	err := c.Bind(&p)
	if err != nil {
		return h.response.BindFailed(err)
	}
	// Get ID from URL and parse to UUID
	ID, err := uuid.Parse(c.Param("id"))
	// Check if UUID is valid
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}
	// Assign ID to model
	p.ID = ID

	err = h.useCase.Update(&p)
	if err != nil {
		return h.response.Error(c, "UseCase Update()", err)
	}
	return c.JSON(h.response.OK(p))
}

func (h *handler) Delete(c echo.Context) error {

	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}
	h.useCase.Delete(ID)
	return c.JSON(h.response.OK(nil))
}

func (h *handler) GetAll(c echo.Context) error {

	// Pagination
	// Page
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1 // Default page
	}
	// Limit
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10 // Default limit
	}

	products, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "UseCase GetAll()", err)
	}
	return c.JSON(h.response.OK(products))
}

func (h *handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}
	p, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "UseCase GetByID()", err)
	}
	return c.JSON(h.response.OK(&p))
}
