package main

import (
	"fmt"
	"sync"
)

func say(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	defer wg.Done()

}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg)
	wg.Wait()

	wg.Add(1)
	go say("Here!!", &wg)

	wg.Wait()
	wg.Add(1)

	go func(text string, wg *sync.WaitGroup) {
		fmt.Print(text)
		defer wg.Done()
	}("Bye!!", &wg)

	wg.Wait()

	var casa_uso int = 10
	fmt.Print(casa_uso)

	//time.Sleep(time.Second * 4)
}
