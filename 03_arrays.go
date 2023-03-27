package main

import "fmt"

/*
Declaring arrays
We need: `name + size + type`:
var myArray [10]int
size must be an integer constant greater than 0;
type — any valid Go data type;
The type [n]T is an array of n values of type T
An array's length is part of its type, so arrays `cannot be resized`.

var myArray1 = [3]int            // will be filled with so called
                                 // zero values, for integers: 0
var myArray2 = [5]int{1,2,3,4,5} // number of values between { } can
                                 // not be larger than size (ofc)
var myArray3 = […]int{1,2,3,4}   // the compiler will count the
                                 // array elements for you
myArray4 := […]string{"ab","abcd","abcde"} // type is inferred from the data stored.
*/

func array_1() {
	println("Array basics")
	println("==================================")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func array_2() {
	/* Define an array of size 4 that stores deployment options */
	var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}
	println("Looping over array")
	println("==================================")
	/* Loop through the deployment options array */
	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}
}

func array_3() {
	/* Define an array and let the compiler count its size */
	deployment_options := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	println("Using range keyword")
	println("==================================")
	for index, option := range deployment_options {
		fmt.Println(index, option)
	}

}

func main() {
	array_1()
	array_2()
	array_3()
}
