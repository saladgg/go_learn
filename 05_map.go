package main

import "fmt"

func main() {
	/*
		- a map is a data structure that stores key-value pairs with a constant look-up time.
		- think of it as a python's dictionary or JS's object.
	*/

	/* Define a map containing the release year of several languages */
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}
	fmt.Println("===========fine releases============")
	finerReleases := map[string]float32{
		"django": 3.2, "drf": 3.4, "python": 3.11, "angular": 11.55,
	}

	for k, v := range finerReleases {
		fmt.Printf("%s fine release %f\n", k, v)
	}
}
