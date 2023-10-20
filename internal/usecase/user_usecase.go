package usecase

import (
	"heroku/internal/delivery/http/model"
	"heroku/internal/entities"
	repo "heroku/internal/repository"
	iError "heroku/pkg/error"
	"net/http"
	"sync"
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

	if userErr != nil {
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, "Error while getting user info")
	}

	if accountsErr != nil {
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, "Error while getting account detail")
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
