package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	// Pointers
	fmt.Printf("%T\n", b)

	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer
	*b = 7
	fmt.Println(a)
}
