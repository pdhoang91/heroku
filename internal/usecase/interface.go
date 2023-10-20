package usecase

import (
	"heroku/internal/delivery/http/model"
	"heroku/internal/entities"
)

type UserUseCase interface {
	GetUserInfo(id int) (*model.UserInfo, error)
}

type BalanceCalculator interface {
	CalculateBalance(accounts []*entities.Account) float32
}
