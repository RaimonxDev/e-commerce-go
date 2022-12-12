package model

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Responses []Response

// MessageResponse standard response for the user
type MessageResponse struct {
	Data     interface{} `json:"data"`
	Errors   Responses   `json:"errors"`
	Messages Responses   `json:"message"`
}
