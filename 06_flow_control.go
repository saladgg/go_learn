package main

import "fmt"

// a function that takes a slice as a parameter and returns a float called avg
func average(marks []float64) (avg float64) {
	total := 0.0

	if len(marks) == 0 {
		avg = 0
	} else {
		for _, v := range marks {
			total += v
		}
		avg = total / float64(len(marks))
	}
	return
}

func main() {
	var scores = [5]float64{90.9, 87.7, 80.9, 96.3, 85.6}
	fmt.Println(average(scores[:]))
}
