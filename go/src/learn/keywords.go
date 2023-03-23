package learn

import "fmt"

// keywords: defer, continue, break
func Kw_ex() {
	// defer -> send as last execution -> useful for closing things
	// multipe defers
	defer fmt.Println("run as ultimate - closing db connection")
	defer fmt.Println("run as penultimate - sending something")
	// continue and break
	for i := 0; i < 10; i++ {
		if i < 2 {
			fmt.Println("ignore at:", i)
			continue
		}
		if i == 4 {
			fmt.Println("break at:", i)
			break
		}
		fmt.Println(i)
	}
}