package main

import "fmt"

// var can be declared outside func also
var age = 24

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var keyword
	var name = "KameR"

	// Shorthand property
	bio := "Designer & Developer"

	// const can't be override
	const isCool = true

	// Can override the type and values also
	var age int32 = 25

	size := 6.2

	// Can assign multiple values
	email, location := "kashyapk62@gmail.com", "India"

	fmt.Println(name, email, age, bio, isCool, size, location)
	fmt.Printf("%T", age)
}
