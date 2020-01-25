/*
	struct grammar
*/

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // create a struct of type vertex
	v2 = Vertex{X: 1}  // Y:0 Y is implicitly assigned 0
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // create a pointer of type *Vertex struct
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
