package usecase

import (
	"heroku/internal/entities"
	"testing"
)

func TestSimpleBalanceCalculator_CalculateBalance(t *testing.T) {
	// Create an instance of SimpleBalanceCalculator
	calculator := &SimpleBalanceCalculator{}

	// Create test cases with various account balances
	testCases := []struct {
		name     string
		accounts []*entities.Account
		expected float32
	}{
		{
			name:     "Empty accounts",
			accounts: []*entities.Account{},
			expected: 0,
		},
		{
			name: "Single account with balance",
			accounts: []*entities.Account{
				{ID: 1, UserId: 1, Name: "A銀行", Balance: 20000},
			},
			expected: 20000,
		},
		{
			name: "Multiple accounts with balances",
			accounts: []*entities.Account{
				{ID: 1, UserId: 1, Name: "A銀行", Balance: 20000},
				{ID: 2, UserId: 1, Name: "B銀行", Balance: 10000},
				{ID: 3, UserId: 1, Name: "C信用金庫", Balance: 15000},
			},
			expected: 45000,
		},
	}

	// Iterate over test cases and run the tests
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := calculator.CalculateBalance(testCase.accounts)

			if result != testCase.expected {
				t.Errorf("Expected balance %f, but got %f for test case %s", testCase.expected, result, testCase.name)
			}
		})
	}
}
