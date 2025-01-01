package services

import (
	"account-management/accounts/models"

	"github.com/sirupsen/logrus"
)

// Accounts slice to store employee data
var Accounts []models.Account
var Transactions []models.Transaction

// AddAccount adds a employee to the list
func AddAccount(id int, name string, balance int) {
	employee := models.Account{ID: id, Name: name, Balance: balance}
	Accounts = append(Accounts, employee)

	logrus.WithFields(logrus.Fields{
		"employee_id":   employee.ID,
		"employee_name": employee.Name,
		"employee_age":  employee.Balance,
	}).Info("Account Added Successfully")

}

// GetAccountByID returns a employee by ID
func GetAccountByID(id int) models.Account {
	for _, employee := range Accounts {
		if employee.ID == id {
			return employee
		}
	}
	return models.Account{}
}

func Deposit(id int, amount int) {
	for i, account := range Accounts {
		if account.ID == id && amount > 0 {
			Accounts[i].Balance += amount
			transaction := models.Transaction{ID: id, Amount: amount}
			Transactions = append(Transactions, transaction)
			break
		}
	}
}

func Withdraw(id int, amount int) {
	for i, account := range Accounts {
		if account.ID == id && (Accounts[i].Balance-amount) > 0 {
			Accounts[i].Balance -= amount
			transaction := models.Transaction{ID: id, Amount: -amount}
			Transactions = append(Transactions, transaction)
			break
		}
	}
}

func GetTransactions(id int) []models.Transaction {
	var transactions []models.Transaction
	for _, transaction := range Transactions {
		if transaction.ID == id {
			transactions = append(transactions, transaction)
		}
	}
	return transactions
}

// GetAllAccounts returns all employees
func GetAllAccounts() []models.Account {
	return Accounts
}
