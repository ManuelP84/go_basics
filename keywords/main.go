package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("End of the main function!")
	fmt.Println("Start main function!")

	// Continue and Break
	for i := 0; i < 10; i++ {
		fmt.Println("Item: ", i)

		if i == 5 {
			fmt.Println("Lets continue!")
			continue
		}

		if i == 7 {
			fmt.Println("Break!!!")
			break
		}
	}
}
