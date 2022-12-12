package model

import "fmt"

type Error struct {
	Code       string
	Err        string
	Who        string
	StatusHTTP int
	Data       interface{}
	ApiMessage string
	userID     string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s, Error: %s, Who:%s, StatusHttp:%d, Data: %v, UserID:%s",
		e.Code, e.Err, e.Who, e.StatusHTTP, e.Data, e.userID)

}

func (e *Error) HasCode() bool {
	return e.Code != ""
}

func (e *Error) HasStatusHttp() bool {
	return e.StatusHTTP > 0
}
func (e *Error) hastData() bool {
	return e.Data != nil
}
