/*
	Interface and implicit implementation
	A type implements an interface by implementing all the methods of an interface. 

	Because no explicit declaration is required, there is no "implements" keyword.

	The implicit interface decouples the definition from the implementation of the interface,
	so that the implementation of the interface can appear in any package without prior preparation.

	Therefore, there is no need to add a new interface name to each implementation, which also encourages clear interface definitions.
*/

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// The method indicates that type T implements interface I, but we don't need to declare this explicitly.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// hello


