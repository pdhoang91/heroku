package repo

import "heroku/internal/entities"

type UserRepository interface {
	FindUserByID(userID int) (*entities.User, error)
	GetUserAccounts(userID int) ([]*entities.Account, error)
	GetAllUser() ([]*entities.User, error)
}
