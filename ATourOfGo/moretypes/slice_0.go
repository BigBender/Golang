/*
	Slice

	Slices provide dynamic size for array elements	
	Type []T represents a slice with element type T

	a[low : high]
	Front closing and back opening section

	a[1 : 4] represents a[1], a[2], a[3]
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

// output: [3 5 7]