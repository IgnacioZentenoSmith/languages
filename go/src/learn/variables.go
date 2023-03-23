package learn

import "fmt"

/* declarations
  var -> can change
	const -> doesnt change once set
*/

func Vars_ex() {
	variables()
	zerovals()
}

func variables() {
	
	fmt.Println("bool: represents a boolean value and is either true or false")
	fmt.Println("Numeric: represents integer types, floating point values, and complex types")
	fmt.Println("string: represents a string value")
	
	//var bool_type bool = true     		// Boolean
  //var int_type int = 5         			// Integer
  //var float32_type float32 = 3.14  	// Floating point number
  //var string_Type string = "Hi!"  	// String

	// var declares 1 or more variables
	//var a = "initial"

	// You can declare multiple variables at once.
	//var b, c int = 1, 2

	// Go will infer the type of initialized variables.
	//var d = true

	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	//var e int

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. 
	// This syntax is only available inside functions.
	//f := "apple"
}

func zerovals() {
/* Zero values */
	fmt.Println("Zero values")
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("int zero val: ", a)
	fmt.Println("float64 zero val: ", b)
	fmt.Println("string zero val: ", c)
	fmt.Println("bool zero val: ", d)
}

/* Boolean type -> "bool"
	var b1 bool = true // typed declaration with initial value => Returns true
	var b2 = true // untyped declaration with initial value => Returns true
	var b3 bool // typed declaration without initial value => Returns false
	b4 := true // untyped declaration with initial value => Returns true
*/

/* Integer type -> "int" -> store a whole number without decimals
Signed integers -> int
	int -> default 32 bits
	int8		8 bits/1 byte	-128 to 127
	int16		16 bits/2 byte	-32768 to 32767
	int32		32 bits/4 byte	-2147483648 to 2147483647
	int64		64 bits/8 byte	-9223372036854775808 to 9223372036854775807

Unsigned integers -> "uint" -> can only store non-negative values:
	uint -> default 32 bits
	uint8		8 bits/1 byte	0 to 255
	uint16	16 bits/2 byte	0 to 65535
	uint32	32 bits/4 byte	0 to 4294967295
	uint64	64 bits/8 byte	0 to 18446744073709551615
*/

/* Float type -> "floatx" -> float data types are used to store positive and negative numbers with a decimal point
	float32	32 bits	-3.4e+38 to 3.4e+38.
	float64	64 bits	-1.7e+308 to +1.7e+308.
*/

/* String type -> "string" -> store a sequence of characters, must be surrounded by double quotes
	var txt1 string = "Hello!" -> Hello!
	var txt2 string -> ""
	txt3 := "World 1" -> World 1
*/