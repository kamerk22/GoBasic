package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greetings method ( value receiver )
func (p Person) greet() string {
	return "Hey! " + p.firstName
}

func main() {

	// Init Person using struct
	person1 := Person{firstName: "KameR", lastName: "MakeR", city: "Ahmedabad", gender: "M", age: 24}
	fmt.Println(person1)

	// Alternative Person using struct
	person2 := Person{"Samantha", "Max", "New York", "F", 22}
	fmt.Println(person2)
	fmt.Println(person2.city)
	person2.age++
	fmt.Println(person2)

	fmt.Println(person2.greet())
}
