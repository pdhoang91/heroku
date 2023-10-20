package model

import "heroku/internal/entities"

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

type UserInfo struct {
	UserID   int                 `json:"user_id"`
	Name     string              `json:"name"`
	Accounts []*entities.Account `json:"accounts"`
	Balance  float32             `json:"balance"`
}
