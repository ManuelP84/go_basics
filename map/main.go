package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["manuel"] = 37
	m["angie"] = 35
	m["juanma"] = 5
	m["maria"] = 1

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	// Find a value
	value, ok := m["manuel"]
	fmt.Println(value, ok)
}
