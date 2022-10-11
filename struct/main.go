package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (employee *Employee) setId(id int) {
	employee.id = id
}

func (employee *Employee) setName(name string) {
	employee.name = name
}

func (employee *Employee) getId() int {
	return employee.id
}

func (employee *Employee) getName() string {
	return employee.name
}

func main() {
	e := Employee{}
	fmt.Printf("employee: %v : %T", e, e)

	fmt.Println()

	e.id = 91185
	e.name = "manuel"
	fmt.Printf("employee: %v : %T", e, e)

	fmt.Println()

	e.setName("Juanma")
	e.setId(101124)
	fmt.Printf("employee: %v : %T", e, e)

	fmt.Println()

	fmt.Println(e.getId())
	fmt.Println(e.getName())
}
