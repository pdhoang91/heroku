package usecase

import "heroku/internal/entities"

type SimpleBalanceCalculator struct{}

// Implement your balance calculation logic here
// Calculate the balance using the accounts provided
// and return the result.
func (sbc *SimpleBalanceCalculator) CalculateBalance(accounts []*entities.Account) float32 {
	if len(accounts) <= 0 {
		return 0
	}
	var balance float32
	for _, acc := range accounts {
		balance += float32(acc.Balance)
	}
	return balance
}
