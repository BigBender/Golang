/*
	method

	Go has no classes. 
	However, you can define methods for struct types.

	A method is a class of functions with special receiver parameters.
	The method receiver is in its own parameter list, between the func keyword and the method name.

	In this example, the Abs method has a receiver named v and type Vertex.

	Method is function

	Remember: a method is just a function with receiver parameters.
	Now this Abs is written as a normal function, and the function has not changed.
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// output
// 5
