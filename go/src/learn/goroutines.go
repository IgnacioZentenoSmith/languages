package learn

import (
	"fmt"
	"sync"
)

func goroutine() {
	// accumulate goroutines
	var wg sync.WaitGroup
	// add goroutines into the group
	wg.Add(2)

	go someFunc("Hello ", &wg)
	go someFunc(" World!", &wg)

	// anonymous goroutine
	/*
	go func(text string) {
		fmt.Println(text)
	}("some text")
	*/

	// Wait for goroutines to end
	wg.Wait()
}

func someFunc(s string, wg *sync.WaitGroup) {
	// execute done as last -> this goroutine ends
	defer wg.Done()
	fmt.Printf(s)
}