/*
	Create slices with make
	Slices can be created with the built-in function make, which is how you create dynamic arrays.

	The make function allocates an array of zero elements and returns a slice that references it:

	a := make ([] int, 5)    // len (a) = 5
	To specify its capacity, pass a third argument to make:

	b := make ([] int, 0, 5) // len (b) = 0, cap (b) = 5

	b = b [: cap (b)]        // len (b) = 5, cap (b) = 5
	b = b [1:]               // len (b) = 4, cap (b) = 4
*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// output
// a len=5 cap=5 [0 0 0 0 0]
// b len=0 cap=5 []
// c len=2 cap=5 [0 0]
// d len=3 cap=3 [0 0 0]



