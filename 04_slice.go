package main

import (
	"fmt"
	"reflect"
)

/*
- An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
flexible view into the elements of an array - a pointer to the array.
- The type []T is a slice with elements of type T.
* Slice consists of:
	- pointer to array
	- the `length` of the segment (e.g. meaning elements 0:5 of the array)
	- slice `capacity` (maximum length it can store)
*/

func slice_from_array() {
	fmt.Println("--------------------------------------")
	fmt.Println("Creating slice from array")
	fmt.Println("--------------------------------------")
	myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // firstly we need an array
	// make a copy of the array. this will result to `easySlice`
	easySlice := myArray[:]
	fmt.Println("copy array to slice\n", easySlice)
	fmt.Println("inspect the type: ", reflect.TypeOf(easySlice).Kind())
	// to make the slice out of
	mySlice1 := myArray[1:5] // slice from 1st up to 5th, 5th not included.
	fmt.Println(mySlice1)
	mySlice2 := mySlice1[0:2] // slicing a slice
	fmt.Println(mySlice2)

}

func slice_2() {
	// slice literal, almost as array with just size omitted
	var letters = []string{"a", "b", "c", "d"}
	nums := []int{1, 2, 3}

	// built-in make function. syntax `make([]T, len, cap) []T`
	// capacility should always be >= len
	var mySlice3 = make([]int, 5, 5)

	// create slice starting from 1-3 with capacity of 4
	mySliceX := letters[1:3:4]
	fmt.Println("sliceX\n", mySliceX)

	/*
		- when the capacity argument is omitted, it defaults to the specified length.
		- fills 5 zeros for int blank spaces for string.
	*/
	var mySlice4 = make([]int, 5)

	// growing a slice
	mySlice4 = append(mySlice4, 34)
	mySlice3 = append(mySlice3, nums...)

	fmt.Println(letters)
	fmt.Println(mySlice3)
	fmt.Println(mySlice4)

}

func main() {
	slice_from_array()
	slice_2()
}
