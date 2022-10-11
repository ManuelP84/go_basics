package main

import (
	"fmt"
	"strings"
)

func main() {
	number := 10

	// Way 1.
	switch mod := number % 2; mod {
	case 0:
		fmt.Printf("%v es par\n", number)
	default:
		fmt.Printf("%v es impar\n", number)
	}

	// Way 2.
	value := 200
	switch {
	case value > 200:
		fmt.Println("Values is greater than 200")
	case value < 200:
		fmt.Println("Value is lower than 200")
	default:
		fmt.Println("Value is 200")
	}

	palabra := "En platzi nunca paramos de aprender"
	a, e, i, o, u := contadorVocales(palabra)
	fmt.Printf("la frase '%s' tiene: \n", palabra)
	fmt.Printf("%d vocales a \n", a)
	fmt.Printf("%d vocales e \n", e)
	fmt.Printf("%d vocales i \n", i)
	fmt.Printf("%d vocales o \n", o)
	fmt.Printf("%d vocales u \n", u)
}
func contadorVocales(palabra string) (int, int, int, int, int) {
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	for _, valor := range palabra {
		variable := strings.ToLower(string(valor))
		switch variable {
		case "a":
			conta++
		case "e":
			conte++
		case "i":
			conti++
		case "o":
			conto++
		case "u":
			contu++
		}
	}
	return conta, conte, conti, conto, contu
}
