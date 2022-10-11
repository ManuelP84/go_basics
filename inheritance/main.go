package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	salary float32
}

func (e FullTimeEmployee) getMessage() string {
	return "This is a full-time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate float32
}

func (e TemporaryEmployee) getMessage() string {
	return "This is a temporary employee"
}

type PrintEmployeeInfo interface {
	getMessage() string
}

func getMessage(printInfo PrintEmployeeInfo) {
	fmt.Println(printInfo.getMessage())
}

func main() {
	// Full-time employee
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Manuel"
	ftEmployee.age = 37
	ftEmployee.id = 91185
	ftEmployee.salary = 3500000.0
	fmt.Printf("FullTimeEmployee: %v \n", ftEmployee)

	// Temporary employee
	tEmployee := TemporaryEmployee{}
	tEmployee.name = "Angie"
	tEmployee.age = 35
	tEmployee.id = 1095
	tEmployee.taxRate = 1000000.0
	fmt.Printf("TemporaryEmployee: %v \n", tEmployee)

	// GetMessage - Interface
	getMessage(ftEmployee)
	getMessage(tEmployee)
}
