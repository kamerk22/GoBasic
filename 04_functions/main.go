package main

import "fmt"

func greetings(s string) string {
	return "Hey! " + s
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println(greetings("KameR"))
	fmt.Println(getSum(2, 3))
}
