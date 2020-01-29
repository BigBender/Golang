/*
	var v Vertex
	ScaleFunc(v, 5) // Compile error!
	ScaleFunc(& v, 5) // OK

	When a method with a pointer as the receiver is called, the receiver can be both a value and a pointer:
	var v Vertex
	v.Scale(5) // OK
	p: = &v
	p.Scale(10) // OK
	For the statement v.Scale(5), even if v is a value rather than a pointer, the method with pointer receiver can be called directly. 
	That is, since the Scale method has a pointer receiver, Go interprets the statement v.Scale (5) as (& v) .Scale (5) for convenience.
*/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// output
// {60 80} &{96 72}


