package main

import (
	"fmt"
	"fake.com/languages/go/src/learn"
)

// run -> go run file.go
// build -> go build go/src/file.go 		
//			 -> ./file

func main() {
	learn.Map_ex()
	// Sprintf -> saves as string
	var message string = fmt.Sprintf("Hello world!")
	// Println -> print with new line
	fmt.Println(message)
	// data type -> use Printf (classic C printf) with %T
	fmt.Printf("%T", message)
}