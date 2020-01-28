/*
	Modify mapping

	Insert or modify elements in map m:
	m [key] = elem

	Get the element:
	elem = m [key]

	Remove element:
	delete (m, key)

	Detect the existence of a key by double assignment:
	elem, ok = m [key]

	If key is in m, ok is true; otherwise, ok is false.

	If key is not in the map, elem is the zero value of the type of the map element.

	Similarly, when a nonexistent key is read from the map, the result is a zero value for the mapped element type.

	Note: If elem or ok have not been declared, you can use short variable declarations:
	elem, ok: = m [key]
*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// output
// The value: 42
// The value: 48
// The value: 0
// The value: 0 Present? false