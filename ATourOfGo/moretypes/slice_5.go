/*
	nil slice
	The zero value of the slice is nil.

	The nil slice is 0 in length and capacity and has no underlying array.
*/
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// output
// [] 0 0
// nil!
