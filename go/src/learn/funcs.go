package learn

import "fmt"

func Funcs_ex() {
	fmt.Println(getInt(123))
	fmt.Println(getString("some str"))
	fmt.Println(getBool(false))
}

// args, returns
func getInt(val int) (r1 string, r2 int) {
	//fmt.Println("int val: ", val)
	//fmt.Printf("%T\n", val)
	fmt.Printf("%T", val)
	return " value: ", val
}

// args, returns
func getString(val string) (r1 string, r2 string) {
	//fmt.Println("str val: ", val)
	//fmt.Printf("%T\n", val)
	fmt.Printf("%T", val)
	return " value: ", val
}

// args, returns
func getBool(val bool) (r1 string, r2 bool) {
	//fmt.Println("bool val: ", val)
	//fmt.Printf("%T\n", val)
	fmt.Printf("%T", val)
	return " value: ", val
}