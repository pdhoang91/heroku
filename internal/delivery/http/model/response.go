package model

type SuccessResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
