package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	//"github.com/google/uuid"
)

func main() {
	// Variable declaration
	const piValue float64 = 3.14
	fmt.Println("Pi: ", piValue)

	// Zero values
	var a float64
	var b bool
	var c int
	var d string
	fmt.Println("a: ", a, "b: ", b, "c: ", c, "d: ", d)

	// Square area
	const base int = 10
	squareArea := base * base
	fmt.Println("area: ", squareArea)

	// Circle area
	const pi = 3.141516
	r := 5.0
	circleArea := pi * math.Pow(float64(r), 2)
	fmt.Println("Circle area: ", circleArea)

	// Math operations
	x := 10
	y := 20

	// Sum
	result := x + y
	fmt.Println("Sum: ", result)

	// Dif
	result = y - x
	fmt.Println("Dif: ", result)

	// Multiplication
	result = x * y
	fmt.Println("Mult: ", result)

	// Division
	result = y / x
	fmt.Println("Div: ", result)

	//Module
	result = y % x
	fmt.Println("Module: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental", x)

	//println(uuid.New().String())

	helloMessage := "Hello"
	worldMessage := "world"

	// Println : Agrega un salto de linea
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nameManuel := "Manuel"
	ageManuel := 37
	fmt.Printf("%s tiene %v años\n", nameManuel, ageManuel)

	// Sprintf
	nameAngi := "Angie"
	ageAngie := 35
	message := fmt.Sprintf("%s tiene %v años", nameAngi, ageAngie)
	fmt.Println(message)

	// Type
	fmt.Printf("helloMessage: %T nameManuel: %T ageManuel: %T\n", helloMessage, nameManuel, ageManuel)

	// Turn a text into a number
	value, err := strconv.Atoi("123w")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Converted value is: %v\n", value)

	argumentsFunction(1, 2, "manuel")
	printMessage("Hello World")
	fmt.Println(returnValue(3, 6))
	value1, value2 := returnTwoValues(5)
	fmt.Println("Value1, Value2: ", value1, value2)
	_, result2 := returnTwoValues(5)
	fmt.Println("Result: ", result2)
}

// Functions
func argumentsFunction(num1, num2 int, message string) {
	fmt.Println(num1, num2, message)
}

func printMessage(message string) {
	fmt.Println(message)
}

func returnValue(num1, num2 int) int {
	return num1 + num2
}

func returnTwoValues(num1 int) (a, b int) {
	return num1, num1 * 3
}
