package main

import "fmt"

func explicit_vars() {
	var a int

	var (
		b bool
		c float32
		d string
	)
	a = 20
	b, c = true, 50.5
	d = "Test string"

	fmt.Println(a, b, c, d)
}

func const_short_vars() {
	/* User specified types */
	const a int32 = 12        // 32-bit integer
	const b float32 = 20.5    // 32-bit float
	var c complex128 = 1 + 4i // 128-bit complex number
	var d uint16 = 14         // 16-bit unsigned integer

	/*
		- Short Variable declaration - you dont need the `var/const` keyword and and variable <type>
		- The type is inferred from the assigned value
	*/
	n := 42              // int
	pi := 3.14           // float64
	x, y := true, false  // bool
	z := "Go is awesome" // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}

func main() {
	explicit_vars()
	const_short_vars()
}
