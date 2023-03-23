package learn

import "fmt"

func For_ex() {

	fmt.Println("-------------For-------------")
	// For
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------For while-------------")
	// For while
	counter := 0
	for counter < 3 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("-------------For forever-------------")
	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 3 {
			break
		}
	}

	fmt.Println("-------------For Range-------------")
	//For Range
	arr := [8]int{0, 1, 4, 6, 10, 9}
	fmt.Println("Arr:", arr)

	for i, j := range arr {
		fmt.Printf("index i: %d has val: %d\n", i, j)
	}
}