package usecase

import (
	"heroku/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock User Repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserByID(userID int) (*entities.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) GetUserAccounts(userID int) ([]*entities.Account, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entities.Account), args.Error(1)
}

// Mock Balance Calculator
type MockBalanceCalculator struct {
	mock.Mock
}

func (m *MockBalanceCalculator) CalculateBalance(accounts []*entities.Account) float64 {
	args := m.Called(accounts)
	return args.Get(0).(float64)
}

func TestIUserUseCase_GetUserInfo(t *testing.T) {
	// Create a new instance of the IUserUseCase with mock dependencies
	mockUserRepo := new(MockUserRepository)
	mockBalanceCalc := new(MockBalanceCalculator)
	useCase := NewUserUseCase(mockUserRepo)

	// Set up the expected values
	expectedUser := &entities.User{ID: 1, Name: "Alice"}
	expectedAccounts := []*entities.Account{
		{ID: 1, UserId: 1, Name: "A銀行", Balance: 20000},
		{ID: 2, UserId: 1, Name: "B銀行", Balance: 10000},
	}
	expectedBalance := float32(30000.0)

	// Mock the UserRepository's FindUserByID method
	mockUserRepo.On("FindUserByID", 1).Return(expectedUser, nil)

	// Mock the UserRepository's GetUserAccounts method
	mockUserRepo.On("GetUserAccounts", 1).Return(expectedAccounts, nil)

	// Mock the BalanceCalculator's CalculateBalance method
	mockBalanceCalc.On("CalculateBalance", expectedAccounts).Return(expectedBalance)

	// Call the GetUserInfo method
	userInfo, err := useCase.GetUserInfo(1)

	// Assert that the expected values are returned
	assert.NoError(t, err)
	assert.NotNil(t, userInfo)
	assert.Equal(t, userInfo.UserID, expectedUser.ID)
	assert.Equal(t, userInfo.Name, expectedUser.Name)
	assert.Equal(t, userInfo.Accounts, expectedAccounts)
	assert.Equal(t, userInfo.Balance, expectedBalance)

	// Assert that the methods of the mocks were called as expected
	mockUserRepo.AssertExpectations(t)
}
