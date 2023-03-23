package learn

import "fmt"

// arrays -> fixed size set at declaration
// slice -> variable size
// lenght -> len(arr) -> returns the length of the slice (the number of elements in the slice)
// capacity -> cap(arr) -> returns the capacity of the slice (the number of elements the slice can grow or shrink to)
func Arr_ex() {
	var array [4] int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := [] int {0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))

	// exclusive index 3
	fmt.Println(slice[:3], "up to 3")
	// inclusive index 2, exclusive index 4
	fmt.Println(slice[2:4], "from 2 to 4")
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 11)
	fmt.Println(slice)

	// Create a Slice From an Array
	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array

	// Append One Slice To Another Slice
	// slice3 = append(slice1, slice2...)

	// Change The Length of a Slice
	// myslice1 = arr1[1:3] // Change length by re-slicing the array
}