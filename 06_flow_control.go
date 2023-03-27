package main

import "fmt"

/*
- A function that takes a slice as a parameter and returns a float called avg
- When you initialize a floating point variable with an initial value without specifying a type explicitly,
- ... the compiler will infer the type as `float64` e.g total := 0.0 will be inferred to type of float64
*/
func calc_avg(marks []float32) (avg float32) {
	if len(marks) == 0 {
		avg = 0
	} else {
		var total float32 = 0.0
		for _, v := range marks {
			total += v
		}
		avg = total / float32(len(marks))
	}
	return
}

func calc_avg2(mks []float64) (avg float64) {
	switch len(mks) {
	case 0:
		avg = 0
	default:
		total := 0.0
		for _, v := range mks {
			total += v
		}
		avg = total / float64(len(mks))
	}
	return
}

/*
- In Go, there is no while keyword. Instead, we use the for keyword followed by a condition and a loop body.
- The only exception is the missing semicolon at the end of the condition
*/

func count_nums(x int) (total int) {
	counter := 1
	switch x {
	case 0:
		total = counter
	default:
		for counter < x {
			counter += counter
		}
		total = counter
	}
	return
}

func main() {
	var scores = [5]float32{90.9, 78.7, 80.9, 96.3, 85.6}
	var scores2 = []float64{62.3, 54.2, 77.4, 98.8, 69.0}
	fmt.Println("====average with if..else=====")
	fmt.Println(calc_avg(scores[:]))
	fmt.Println("====average with switch case===")
	fmt.Println(calc_avg2(scores2))
	fmt.Println("====while condition is true===")
	fmt.Println(count_nums(20))

}
