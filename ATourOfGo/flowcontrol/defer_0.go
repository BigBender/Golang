package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// The defer statement will defer execution of the function until the outer function returns
// It's remind me of the property, "defer", of the HTML label -- <script> 