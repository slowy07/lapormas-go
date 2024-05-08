package model

type WebResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type WebErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
