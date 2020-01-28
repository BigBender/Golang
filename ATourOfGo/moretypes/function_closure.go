/*
	Function closure
	Go functions can be a closure.
	A closure is a function value that references a variable outside its function body.
	The function can access and assign values to the variables it references, in other words, the function is "bound" to these variables.

	For example, the function adder returns a closure. Each closure is bound to its own sum variable.
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/* 
	output
	0 0
	1 -2
	3 -6
	6 -12
	10 -20
	15 -30
	21 -42
	28 -56
	36 -72
	45 -90
*/
