package main

import (
	accounts "account-management/accounts"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// AppLogger is a globally accessible logger
var AppLogger *logrus.Logger

func init() {
	AppLogger = logrus.New()
	AppLogger.SetFormatter(&logrus.JSONFormatter{}) // Set the log format to JSON
	AppLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	}) // Set the log format to text
	AppLogger.SetOutput(os.Stdout)        // Set the output to stdout
	AppLogger.SetLevel(logrus.DebugLevel) // Set the log level to debug
}

func main() {
	var choice int
	for {
		fmt.Println("\n=== Banking System Menu ===")
		fmt.Println("1. Add New Account")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. View Account")
		fmt.Println("5. View Transactions")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var name string
			var balance int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Initial Balance: ")
			fmt.Scanln(&balance)
			accounts.AddAccount(id, name, balance)

		case 2:
			var id int
			var amount int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Amount to Deposit: ")
			fmt.Scanln(&amount)
			accounts.Deposit(id, amount)

		case 3:
			var id int
			var amount int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Amount to Withdraw: ")
			fmt.Scanln(&amount)
			accounts.Withdraw(id, amount)

		case 4:
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			account := accounts.GetAccountByID(id)
			fmt.Printf("Account Details: %v\n", account)

		case 5:
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&id)
			transactions := accounts.GetTransactions(id)
			fmt.Println("Transaction History:")
			for _, t := range transactions {
				fmt.Printf("%v\n", t)
			}

		case 6:
			fmt.Println("Thank you for using our banking system!")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
