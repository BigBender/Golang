/*
	You can declare methods for pointer receivers.

	This means that for a certain type T, the receiver type can use the * T grammar. (Also, T cannot be a pointer like * int.)
	For example, the Scale method is defined here for * Vertex.

	Pointer receiver methods can modify the value pointed to by the receiver (just like Scale does here). 
	Because methods often need to modify their receivers, pointer receivers are more commonly used than value receivers.
	Try removing * from line 16 of the Scale function declaration and observe how the program's behavior changes.

	If a value receiver is used, the Scale method operates on a copy of the original Vertex value. 
	(The same is true for the other parameters of the function.) 

	The Scale method must use a pointer receiver to change the value of Vertex declared in the main function.
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

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// output
// 50
