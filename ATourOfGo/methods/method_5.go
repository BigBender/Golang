/*
	Functions that accept a value as a parameter must accept a value of the specified type:

	var v Vertex
	fmt.Println (AbsFunc (v)) // OK
	fmt.Println (AbsFunc (&v)) // Compile error!

	When a "method" with a value is called, the receiver can be both a value and a pointer:

	var v Vertex
	fmt.Println (v.Abs ()) // OK
	p := &v
	fmt.Println (p.Abs ()) // OK

	In this case, the method call p.Abs() is interpreted as (*p).Abs()
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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// *p as p
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

// output
// 5
// 5
// 5
// 5



