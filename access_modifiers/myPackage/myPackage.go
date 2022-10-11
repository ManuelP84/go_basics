package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

// Person private
type Person struct {
	Name string
	Id   string
	Age  int
}

// Print a message
func PrintMessage(message string) {
	fmt.Println(message)
}
