package main

import (
	"GoBasic/03_packages/strutils"
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(91283))

	// Custom package method
	fmt.Println(strutils.Reverse("dlrow olleh"))
}
