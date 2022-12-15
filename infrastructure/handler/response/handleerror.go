package response

import (
	"github.com/RaimonxDev/e-commerce-go.git/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	// custom error
	e, ok := err.(*model.Error)
	if ok {
		_ = c.JSON(getResponseError(e))
		return
	}

	// check echo error
	if echoErr, ok := err.(*echo.HTTPError); ok {
		msg, ok := echoErr.Message.(string)
		if !ok {
			msg = "unknown error"
		}

		_ = c.JSON(echoErr.Code, model.MessageResponse{
			Errors: model.Responses{
				{Code: UnexpectedError, Message: msg},
			},
		})
		return
	}

	// if the handler not returns a "model.Error" then it returns a generic error JSON response
	_ = c.JSON(http.StatusInternalServerError, model.MessageResponse{
		Errors: model.Responses{
			{Code: UnexpectedError, Message: "unknown error"},
		},
	})
}

// getResponseError returns the status code and a Response
func getResponseError(err *model.Error) (int, model.MessageResponse) {
	outputStatus := 0
	outputResponse := model.MessageResponse{}
	if !err.HasCode() {
		err.Code = UnexpectedError
	}

	if err.HasData() {
		outputResponse.Data = err.Data
	}

	if !err.HasStatusHttp() {
		err.StatusHTTP = http.StatusInternalServerError
	}

	outputStatus = err.StatusHTTP
	outputResponse.Errors = model.Responses{model.Response{
		Code:    err.Code,
		Message: err.ApiMessage,
	}}

	return outputStatus, outputResponse
}