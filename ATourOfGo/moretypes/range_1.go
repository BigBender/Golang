/*
	You can assign a subscript or value to _ to ignore it.

	for i, _: = range pow
	for _, value: = range pow
	If you only need the index, ignore the second variable.

	for i: = range pow
*/

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// output
// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512