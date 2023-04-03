package main

import "fmt"

/*
- A struct (short for structure) is used to create a collection of fields of different data types, into a single variable.
- To declare a structure in Go, use the type and struct keywords. e.g type var_name struct {...}
- While arrays are used to store multiple values of the same data type into a single variable,
- structs are used to store multiple values of different data types into a single variable.
*/

type Person struct {
	name   string
	age    int
	job    string
	salary float32
}

func main() {
	var person_1 Person

	person_1.name = "jane"
	person_1.age = 23
	person_1.job = "Cashier"
	person_1.salary = 9700.00

	// Access and print Pers1 info
	fmt.Println("Name: ", person_1.name)
	fmt.Println("Age: ", person_1.age)
	fmt.Println("Job: ", person_1.job)
	fmt.Println("Salary: ", person_1.salary)
}
