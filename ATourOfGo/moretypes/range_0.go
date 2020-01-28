/*
	Range
	The range form of a "for" loop can traverse slices or maps.
	When using a "for" loop to traverse a slice, two values are returned per iteration
	The first value is the index of the current element
	and the second value is a copy of the element corresponding to the index.
*/

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// output
// 2**0 = 1
// 2**1 = 2
// 2**2 = 4
// 2**3 = 8
// 2**4 = 16
// 2**5 = 32
// 2**6 = 64
// 2**7 = 128