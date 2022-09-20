package data

import "github.com/claravelita/training-golang-mnc/dtos"

func EmployeesData() []dtos.Employee {
	employees := []dtos.Employee{
		{
			ID:       1,
			Name:     "Airell",
			Age:      23,
			Division: "IT",
		},
		{
			ID:       2,
			Name:     "Nanda",
			Age:      23,
			Division: "Finance",
		},
		{
			ID:       3,
			Name:     "Mailo",
			Age:      20,
			Division: "IT",
		},
	}
	return employees
}
