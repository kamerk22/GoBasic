package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["Bruce"] = "brucew@yne.com"
	emails["KameR"] = "k@mer.com"
	emails["MakeR"] = "m@ker.com"

	fmt.Println(emails)
	fmt.Println(emails["Bruce"])

	// Delete from map
	delete(emails, "MakeR")
	fmt.Println(emails)

	// Declare and add values
	ages := map[string]int{"Bruce": 26, "KameR": 24, "Maker": 26}
	fmt.Println(ages)
}
