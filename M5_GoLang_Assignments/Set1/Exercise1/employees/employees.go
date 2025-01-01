package employees

import (
	"employee-management/employees/models"
	"employee-management/employees/services"
)

// AddEmployee adds a employee to the list
func AddEmployee(id int, name string, age int, department string) {
	services.AddEmployee(id, name, age, department)
}

// GetEmployeeByID returns a employee by ID
func GetEmployeeByID(id int) models.Employee {
	return services.GetEmployeeByID(id)
}

// GetEmployeebyDepartment returns a employee by department
func GetEmployeebyDepartment(department string) []models.Employee {
	return services.GetEmployeebyDepartment(department)
}

// GetAllEmployees returns all employees
func GetAllEmployees() []models.Employee {
	return services.GetAllEmployees()
}
