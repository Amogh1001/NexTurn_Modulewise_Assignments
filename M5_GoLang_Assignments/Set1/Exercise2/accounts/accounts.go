package employees

import (
	"account-management/accounts/models"
	"account-management/accounts/services"
)

// AddAccount adds a employee to the list
func AddAccount(id int, name string, age int) {
	services.AddAccount(id, name, age)
}

// GetAccountByID returns a employee by ID
func GetAccountByID(id int) models.Account {
	return services.GetAccountByID(id)
}

// Deposit adds an amount to the account
func Deposit(id int, amount int) {
	services.Deposit(id, amount)
}

// Withdraw subtracts an amount from the account
func Withdraw(id int, amount int) {
	services.Withdraw(id, amount)
}

// GetTransactions returns all transactions for an account
func GetTransactions(id int) []models.Transaction {
	return services.GetTransactions(id)
}

// GetAllAccounts returns all employees
func GetAllAccounts() []models.Account {
	return services.GetAllAccounts()
}
