package main

import (
	"employee-management/employees"
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

	AppLogger.Info("Starting the application...")

	// Using the employees package
	employees.AddEmployee(1, "John Doe", 30, "IT")
	employees.AddEmployee(2, "Jane Doe", 25, "HR")

	// Get all employees
	fmt.Println("All employees:", employees.GetAllEmployees())

	// Get employee by ID
	fmt.Println("Employee by ID:", employees.GetEmployeeByID(1))

	// Get employee by department
	fmt.Println("Employee by Department:", employees.GetEmployeebyDepartment("HR"))

}
