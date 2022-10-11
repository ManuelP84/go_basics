package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 20; i++ {
		fmt.Println("Item: ", i)
	}

	// For while
	counter := 0
	for counter < 20 {
		fmt.Println("Item: ", counter)
		counter++
	}

	// For forever
	countForever := 0
	for {
		fmt.Println(countForever)
		countForever++
		if countForever > 100 {
			break
		}
	}

	// For range
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range numbers {
		if i%2 == 0 && i != 0 {
			fmt.Printf("%d is even\n", i)
		}
	}
}
