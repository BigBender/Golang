package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// Execute a simple statement before a conditional expression
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
// difference between "if" in C and GoLang