package main

import "fmt"

func main() {

	// Arrays
	var cars [2]string

	// Assign values
	cars[0] = "Aston Martin"
	cars[1] = "Cadillac"

	fmt.Println(cars)
	fmt.Println(cars[0])

	// Declare and Assign
	movies := [2]string{"Interstellar", "The Dark Knight"}

	fmt.Println(movies)

	// Slices
	series := []string{"Black Mirror", "Narcos", "Breaking Bad"}

	fmt.Println(series)
	fmt.Println(len(series))
	fmt.Println(series[1:3])

}
