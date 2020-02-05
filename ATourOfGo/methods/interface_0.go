/*
	Interface

	An interface type is a collection defined by a set of method signatures.

	An interface type variable can hold any value that implements these methods.

	Note: There is an error in line 22 of the sample code. 

	Because the Abs method is only defined for *Vertex (pointer type), Vertex (value type) does not implement Abser.
*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// The "v" in the following line is a type of Vertex but not *Vertex
	// so it doesn't implement Abser。
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// output
// 5
