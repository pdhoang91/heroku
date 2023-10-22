package usecase

import (
	"fmt"
	"heroku/internal/delivery/http/model"
	"heroku/internal/entities"
	repo "heroku/internal/repository"
	iError "heroku/pkg/error"
	"net/http"
	"sync"

	"github.com/labstack/gommon/log"
)

type IUserUseCase struct {
	UserRepository repo.UserRepository
	BalanceCalc    BalanceCalculator
}

func NewUserUseCase(userRepo repo.UserRepository) UserUseCase {
	return &IUserUseCase{
		UserRepository: userRepo,
		BalanceCalc:    &SimpleBalanceCalculator{}, // Provide the implementation here to adapt business
	}
}

func (uc *IUserUseCase) GetUserInfo(userID int) (*model.UserInfo, error) {
	var user *entities.User
	var accounts []*entities.Account
	var userErr, accountsErr error

	var wg sync.WaitGroup

	// Retrieve user information concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		user, userErr = uc.UserRepository.FindUserByID(userID)
	}()

	// Retrieve account information concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, accountsErr = uc.UserRepository.GetUserAccounts(userID)
	}()

	wg.Wait() // Wait for both operations to complete

	if userErr != nil || user == nil {
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, fmt.Sprintf("User [%d] not found", userID))
	}

	if accountsErr != nil || accounts == nil {
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, fmt.Sprintf("Account of user [%d] not found", userID))
	}

	// Calculate the balance using the strategy pattern
	balance := uc.BalanceCalc.CalculateBalance(accounts)

	// Build and return the user information
	userInfo := model.UserInfo{
		UserID:   user.ID,
		Name:     user.Name,
		Accounts: accounts,
		Balance:  balance,
	}

	return &userInfo, nil
}

func (uc *IUserUseCase) GetAllUserInfo() ([]*model.UserInfo, error) {

	var allUserInfo []*model.UserInfo
	var users []*entities.User
	var accounts []*entities.Account

	//var wg sync.WaitGroup

	users, err := uc.UserRepository.GetAllUser()

	if err != nil {
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, "User not found")
	}

	for _, user := range users {
		accounts, err = uc.UserRepository.GetUserAccounts(user.ID)
		if err != nil || accounts == nil {
			log.Print(fmt.Sprintf("Account of user [%d] not found", user.ID))
		}
		// Calculate the balance using the strategy pattern
		balance := uc.BalanceCalc.CalculateBalance(accounts)

		// Build and return the user information
		userInfo := model.UserInfo{
			UserID:   user.ID,
			Name:     user.Name,
			Accounts: accounts,
			Balance:  balance,
		}

		allUserInfo = append(allUserInfo, &userInfo)
	}

	return allUserInfo, nil
}
