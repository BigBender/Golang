/*
	When slicing, you can use its default behavior to ignore upper and lower bounds.

	The default value of the slice lower bound is 0, and the upper bound is the length of the slice.

	For array
	var a [10] int

	The following slices are equivalent:
	a [0:10]
	a [: 10]
	a [0:]
	a [:]
*/
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// output
// [3 5 7]
// [3 5]
// [5]