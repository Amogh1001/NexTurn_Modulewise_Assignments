package services

import (
	"employee-management/employees/models"

	"github.com/sirupsen/logrus"
)

// Employees slice to store employee data
var Employees []models.Employee

// AddEmployee adds a employee to the list
func AddEmployee(id int, name string, age int, department string) {
	employee := models.Employee{ID: id, Name: name, Age: age, Department: department}
	Employees = append(Employees, employee)

	logrus.WithFields(logrus.Fields{
		"employee_id":   employee.ID,
		"employee_name": employee.Name,
		"employee_age":  employee.Age,
		"department":    employee.Department,
	}).Info("Employee Added Successfully")

}

// GetEmployeeByID returns a employee by ID
func GetEmployeeByID(id int) models.Employee {
	for _, employee := range Employees {
		if employee.ID == id {
			return employee
		}
	}
	return models.Employee{}
}

func GetEmployeebyDepartment(department string) []models.Employee {
	var employees []models.Employee
	for _, employee := range Employees {
		if employee.Department == department {
			employees = append(employees, employee)
		}
	}
	return employees
}

// GetAllEmployees returns all employees
func GetAllEmployees() []models.Employee {
	return Employees
}
