/*
	You can also declare methods for non-struct types.

	In this example, we see a numeric type MyFloat with an Abs method.
	You can only declare methods for receivers of types defined in the same package. 
	You cannot declare methods for receivers of types defined in other packages, including built-in types such as int.

	(Annotation: The receiver's type definition and method declaration must be in the same package; methods cannot be declared for built-in types.)
*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// output
// 1.4142135623730951


