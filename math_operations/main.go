package main

import (
	"fmt"
	"math"
)

// Sum 2 int numbers
func sumNumbers(num1, num2 int) int {
	return num1 + num2
}

// Substraction of 2 int numbers
func subNumbers(num1, num2 int) int {
	return num1 - num2

}

// Multiplication of 2 int numbers
func multNumbers(num1, num2 int) int {
	return num1 * num2
}

// Division from 2 numbers
func divisionNumbers(num1, num2 float32) float32 {
	return num1 / num2
}

// Circle area
func circleArea(r float64) (radius, area float64) {
	return r, math.Pi * math.Pow(r, 2)
}

// Rectangle area
func rectangleArea(base, hight float64) (b, h, a float64) {
	return base, h, base * hight
}

func main() {
	fmt.Println(sumNumbers(4, 5))
	fmt.Println(subNumbers(4, 6))
	fmt.Println(multNumbers(4, 5))
	fmt.Println(divisionNumbers(4, 5))
	radius, area := circleArea(5)
	fmt.Printf("The radius is: %v and the area is: %v \n", radius, area)
	_, _, rectangleArea := rectangleArea(6, 3)
	fmt.Printf("The area of the rectangle is: %v \n", rectangleArea)

}
