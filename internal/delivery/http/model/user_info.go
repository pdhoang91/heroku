package model

import "heroku/internal/entities"

type UserInfo struct {
	UserID   int                 `json:"user_id"`
	Name     string              `json:"name"`
	Accounts []*entities.Account `json:"accounts"`
	Balance  float32             `json:"balance"`
}
