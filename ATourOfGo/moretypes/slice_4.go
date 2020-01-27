/*
	Slice length and capacity

	The length of a slice is the number of elements it contains.
	The capacity of a slice is counted from its first element to the end of its underlying array element.

	The length and capacity of the slice s can be obtained by the expressions len (s) and cap (s).
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// take a slice to have a length of 0
	s = s[:0]
	printSlice(s)

	// expand its length
	s = s[:4]
	printSlice(s)

	// discard the first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// output
// len=6 cap=6 [2 3 5 7 11 13]
// len=0 cap=6 []
// len=4 cap=6 [2 3 5 7]
// len=2 cap=4 [5 7]
