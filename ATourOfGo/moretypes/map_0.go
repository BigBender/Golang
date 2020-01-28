/*
	Map
	Maps map keys to values.

	The mapped zero value is nil. 
	The nil map has neither keys nor keys.

	The make function returns a map of the given type and initializes it for use.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// output
// {40.68433 -74.39967}