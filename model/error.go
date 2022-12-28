package model

import (
	"errors"
	"fmt"
)

var (
	ErrorInvalidID = errors.New("invalid id")
)

type Error struct {
	Code       string
	Err        error
	Who        string
	StatusHTTP int
	Data       interface{}
	ApiMessage string
	UserID     string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s, Error: %s, Who:%s, StatusHttp:%d, Data: %v, UserID:%s",
		e.Code, e.Err, e.Who, e.StatusHTTP, e.Data, e.UserID)

}
func NewError() Error {
	return Error{}
}

func (e *Error) HasCode() bool {
	return e.Code != ""
}

func (e *Error) HasStatusHttp() bool {
	return e.StatusHTTP > 0
}
func (e *Error) HasData() bool {
	return e.Data != nil
}
