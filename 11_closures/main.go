package main

import "fmt"

func substractor() func(int) int {
	final := 0
	return func(x int) int {
		final -= x
		return final
	}
}

func main() {
	result := substractor()
	for i := 0; i < 10; i++ {
		fmt.Println(result(i))
	}
}
