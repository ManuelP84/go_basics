package main

import (
	"fmt"

	"github.com/ManuelP84/go_packages/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrary"
	fmt.Println(myCar)

	var person mypackage.Person
	person.Age = 37
	person.Id = "1234"
	person.Name = "manuel"
	fmt.Println(person)

	mypackage.PrintMessage("Hello!!")
}
