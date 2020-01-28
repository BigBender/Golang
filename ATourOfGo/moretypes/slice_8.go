/*
	Append elements to a slice
	Appending new elements to a slice is a common operation, for which Go provides a built-in append function. 
	The documentation of the built-in function has a detailed introduction to this function.

	func append (s [] T, vs ... T) [] T
	The first parameter s of append is a slice with element type T,
	and the remaining values of type T are appended to the end of the slice.
	
	The result of append is a slice containing all the elements of the original slice plus the newly added elements.

	When the underlying array of s is too small to hold all the given values, it allocates a larger array. 
	The returned slice will point to this newly allocated array.
*/

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// add a empty slice
	s = append(s, 0)
	printSlice(s)

	// this slice will grow on demand
	s = append(s, 1)
	printSlice(s)

	// it can add multiple elements at once
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// output
// len=0 cap=0 []
// len=1 cap=2 [0]
// len=2 cap=2 [0 1]
// len=5 cap=8 [0 1 2 3 4]
