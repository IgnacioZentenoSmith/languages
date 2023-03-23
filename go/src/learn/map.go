package learn

import "fmt"

/* Maps in Go
	Each element in a map is a key:value pair.
	A map is an unordered and changeable collection that does not allow duplicates.
	The length of a map is the number of its elements. You can find it using the len() function.
	The default value of a map is nil.
	Maps hold references to an underlying hash table.
*/
func Map_ex() {
	/* Creating Maps Using var and :=
		var a = map [KeyType] ValueType {key1:value1, key2:value2, ...}
		b := map [KeyType] ValueType {key1:value1, key2:value2, ...}
	*/
	
	var a = map [string] string {"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map [string] int {"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)

	//Knowing if a value exist
	//If we display a["brand"] it will show the value
	value := a["brand"]
	fmt.Println(value)

	//But what if it doesn't exist
	value = a["aasd"]
	fmt.Println(value) //It will show zero

	value, ok := a["aasd"] //If we want to know if it exists, it must be done with the second returing value
	fmt.Println(value , ok)
}