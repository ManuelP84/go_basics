package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Contructor way 1
	employee_one := Employee{}
	fmt.Printf("Employee 1: %v", employee_one)
	fmt.Println()

	// Contructor way 2
	employee_two := Employee{
		id:       1,
		name:     "manuel",
		vacation: true,
	}
	fmt.Printf("Employee 2: %v", employee_two)
	fmt.Println()

	// Contructor way 3
	employee_three := new(Employee)
	fmt.Printf("Employee 3 (pointer): %v", employee_three)
	fmt.Println()
	fmt.Printf("Employee 3 (value): %v", *employee_three)

	fmt.Println()

	employee_three.id = 91185
	employee_three.name = "manuel P"
	fmt.Printf("Employee 3 (pointer): %v", employee_three)
	fmt.Println()
	fmt.Printf("Employee 3 (value): %v", *employee_three)

	fmt.Println()

	// Contructor way 4
	employee_four := NewEmployee(101124, "juanma", true)
	fmt.Printf("Employee 4 : %v", *employee_four)
}
