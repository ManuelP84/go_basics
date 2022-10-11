package main

import (
	"fmt"

	"github.com/ManuelP84/go_pointers/pc"
)

func main() {

	fmt.Println()

	a := 50
	b := &a
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("*b: ", *b)
	fmt.Println("&a : ", &a)

	fmt.Println()

	*b = 500
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("*b: ", *b)
	fmt.Println("&a : ", &a)

	fmt.Println()

	myPc := pc.Pc{Ram: 16, Disk: 200, Brand: "msi"}

	fmt.Println(myPc)

	myPc.DuplicateRam()
	fmt.Println(myPc)

	myPc.DuplicateRam()
	fmt.Println(myPc)
}
