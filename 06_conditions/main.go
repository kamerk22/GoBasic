package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	x := random(1, 6)
	y := random(1, 10)

	// if else
	if x <= y {
		fmt.Printf("%d is less than or equal %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	availableColors := []string{"red", "green", "no"}
	n := rand.Int() % len(availableColors)
	color := availableColors[n]
	// switch
	switch color {
	case "red":
		fmt.Println("Color is Red")
		break
	case "green":
		fmt.Println("Color is Green")
		break
	default:
		fmt.Println("Man I don't know any color like that")
	}

}
