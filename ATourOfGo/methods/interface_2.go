/*
	Interface value

	Interfaces are also values. They can be passed like other values.

	Interface values can be used as parameters or return values for functions.

	Internally, interface values can be viewed as tuples containing values and specific types: (value, type)

	The interface value holds a specific value of a specific underlying type.

	When an interface value calls a method, a method with the same name as its underlying type is executed.
*/

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output
// (&{Hello}, *main.T)
// Hello
// (3.141592653589793, main.F)
// 3.141592653589793


