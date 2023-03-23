package learn

import (
	"fmt"
	"log"
	"strconv"
)

func If_ex() {

	if true {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Bye World")
	}

	//Converting strings to integers
	value, err := strconv.Atoi("23") // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

	if err != nil {
		log.Fatal(err) //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	fmt.Println(value)

	//Even or odd
	numb := 10
	if numb%2 != 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	//login
	user := "Pedro"
	password := "12345"

	if user == "Pedro" && password == "12345" {
		fmt.Println("Logged in")
	} else if user == "Pedro" {
		fmt.Println("Password incorrect")
	} else if password == "12345" {
		fmt.Println("User incorrect")
	} else {
		fmt.Println("Credencials incorrect")
	}

}