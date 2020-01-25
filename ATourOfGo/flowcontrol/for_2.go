package main

import "fmt"

func main() {
	sum := 1
	// Same as "while" in c
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
// "while" in GoLang