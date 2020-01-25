package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Deferred function calls are pushed into a Stack.
// When the outer function returns, the deferred functions are called in the order of last in, first out.