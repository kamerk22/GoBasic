package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Define person struct shorthand
type PersonShortHand struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greetings method ( value receiver )
func (p Person) greet() string {
	return "Hey! My name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age) + " year old."
}

// Has Birthday ( pointer receiver )
func (p *Person) hasBirthday() {
	p.age++
}

// Get Married ( pointer receiver )
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
func main() {

	// Init Person using struct
	person1 := Person{firstName: "KameR", lastName: "MakeR", city: "Ahmadabad", gender: "M", age: 24}
	fmt.Println(person1)

	// Alternative Person using struct
	person2 := Person{"Samantha", "Max", "New York", "F", 22}

	// Shorthand Person using struct
	person3 := PersonShortHand{"Eva", "Max", "New York", "F", 18}
	fmt.Println(person3)
	fmt.Println(person2)
	fmt.Println(person2.city)
	person2.age++
	fmt.Println(person2)

	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Max")
	fmt.Println(person1.greet())

	fmt.Println(person2.greet())
	person2.getMarried("Maker")
	fmt.Println(person2.greet())

}
