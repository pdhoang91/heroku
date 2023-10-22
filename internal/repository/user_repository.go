package repo

import (
	"encoding/json"
	"fmt"
	"heroku/config"
	"heroku/internal/entities"
	iError "heroku/pkg/error"
	"io"
	"net/http"
)

type IUserRepo struct {
}

func NewUserRepository() UserRepository {
	return &IUserRepo{}
}

func sendHTTPRequest(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return iError.NewErrorHandler(response.StatusCode, "Not Found")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	return nil
}

func (ur *IUserRepo) FindUserByID(userID int) (*entities.User, error) {
	url := fmt.Sprintf("%s/users/%d", config.BaseAPIURL, userID)
	var user *entities.User
	err := sendHTTPRequest(url, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *IUserRepo) GetUserAccounts(userID int) ([]*entities.Account, error) {
	url := fmt.Sprintf("%s/users/%d/accounts", config.BaseAPIURL, userID)
	var accounts []*entities.Account
	err := sendHTTPRequest(url, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (ur *IUserRepo) GetAllUser() ([]*entities.User, error) {
	url := fmt.Sprintf("%s/users", config.BaseAPIURL)
	var users []*entities.User
	err := sendHTTPRequest(url, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
