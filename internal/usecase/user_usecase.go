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

var numberOfWorker = 5

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

	users, err := uc.UserRepository.GetAllUser()
	if err != nil || len(users) == 0 {
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, "User not found")
	}

	userQueue := make(chan *entities.User, len(users))
	userInfoChan := make(chan *model.UserInfo, len(users))
	errChan := make(chan error, len(users))

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for _, user := range users {
		userQueue <- user
	}

	close(userQueue)

	for i := 0; i < numberOfWorker; i++ {
		wg.Add(1)
		go uc.worker(&wg, userQueue, userInfoChan, errChan)

	}

	wg.Wait()
	close(userInfoChan)
	close(errChan)

	// Collect user info and errors
	for userInfo := range userInfoChan {
		allUserInfo = append(allUserInfo, userInfo)
	}

	// Handle errors if any
	for err := range errChan {
		log.Print("Error while getting account info:", err)
	}

	return allUserInfo, nil
}

func (uc *IUserUseCase) worker(wg *sync.WaitGroup, userQueue <-chan *entities.User, userInfoChan chan<- *model.UserInfo, errChan chan<- error) {
	defer wg.Done()
	for user := range userQueue {
		userInfo, err := uc.processUser(user)
		if err != nil {
			errChan <- err
		} else {
			userInfoChan <- userInfo
		}
	}
}

func (uc *IUserUseCase) processUser(user *entities.User) (*model.UserInfo, error) {
	accounts, err := uc.UserRepository.GetUserAccounts(user.ID)
	if err != nil || accounts == nil {
		log.Printf("Account of user [%d] not found", user.ID)
		return nil, iError.NewErrorHandler(http.StatusInternalServerError, fmt.Sprintf("Account of user [%d] not found", user.ID))
	}

	// Calculate the balance using the strategy pattern
	balance := uc.BalanceCalc.CalculateBalance(accounts)

	// Build user information
	userInfo := &model.UserInfo{
		UserID:   user.ID,
		Name:     user.Name,
		Accounts: accounts,
		Balance:  balance,
	}

	return userInfo, nil
}
