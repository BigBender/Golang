/*
	Now we are going to rewrite the Abs and Scale methods as functions.

	Similarly, we first try to remove the 16th *. 

	Can you see why the behavior of the program has changed? 
	What should I do to make the example compile successfully?
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

// output 
// 50