package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Varaible assignment

	var numberOne int = 9
	numberTwo := 10

	fmt.Println(numberOne)
	fmt.Println(numberTwo)

	myValue, err := strconv.ParseInt("3", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Map

	m := make(map[string]int)
	m["manuel"] = 37
	m["juanma"] = 5
	m["angie"] = 35
	m["mariap"] = 1

	for member := range m {
		fmt.Print(member, "\n")
	}

	// Slice
	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Println(index, ": ", value)
	}

	s = append(s, 18)

	for index, value := range s {
		fmt.Println(index, ": ", value)
	}
	c := make(chan int)
	go doSomething(c)
	fmt.Println(<-c)

	//time.Sleep(4 * time.Second)
}

func doSomething(c chan int) {
	fmt.Println("Done")
	c <- 1

}
