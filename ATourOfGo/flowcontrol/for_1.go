package main

import "fmt"

func main() {
	sum := 1
	for /* optional */; sum < 1000; /* optional */{
		sum += sum
	}
	fmt.Println(sum)
}
// Init and post statements are optional