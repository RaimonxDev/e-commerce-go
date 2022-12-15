package response

import (
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnexpectedError = "unexpected_error"
	AuthError       = "authorization_error"
)

type API struct{}

func New() API {
	return API{}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: "OK", Message: "success"}},
	}
}
func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: "CREATED", Message: "success"}},
	}
}

func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: "UPDATED", Message: "success"}},
	}
}

func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: "DELETED", Message: "success"}},
	}
}

func (a API) BindFailed(err error) *model.Error {
	e := model.NewError()
	e.StatusHTTP = http.StatusBadRequest
	e.Code = "BIND_FAILED"
	e.Who = "c.Bind()"
	e.Err = err
	// e.Error() is a method of the Error struct in model/error.go
	log.Warnf("Bind failed: %v", e.Error())
	return &e
}

func (a API) Error(c echo.Context, who string, err error) *model.Error {
	e := model.NewError()
	e.StatusHTTP = http.StatusInternalServerError
	e.Code = "INTERNAL_SERVER_ERROR"
	e.ApiMessage = "Internal server error"
	e.Who = who
	e.Err = err

	// Get the user that is causing the error
	userID, ok := c.Get("userID").(uuid.UUID)

	if !ok {
		log.Errorf("Error: %v", e.Error())
	}
	e.UserID = userID.String()
	log.Errorf("Internal server error: %v", e.Error())
	return &e
}
