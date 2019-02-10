package main

import "fmt"

func main() {
	ids := []int{33, 34, 35, 67, 89, 22}

	// Loop through
	for i, id := range ids {
		fmt.Printf("%d - id: %d\n", i, id)
	}

	// Not using Index
	for _, id := range ids {
		fmt.Printf("id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum is %d\n", sum)

	// Range with map
	ages := map[string]int{"Bruce": 26, "KameR": 24, "Maker": 26}
	for k, v := range ages {
		fmt.Printf("%s: %d\n", k, v)
	}
}
