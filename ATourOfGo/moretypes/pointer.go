/*
	Go has got pointer. Pointer holds the memory address of the value.
	Type *T is a pointer points to type T. Its zero value is nil
	
	var p *int

	Operator "&", it creates pointer points to its operand

	i := 42
	p = &i

	Operator "*" gets the underlaying value pointed to by the pointer

	fmt.Println(*p) // Get i by pointer p
	*p = 21         // Set i by pointer p
	
	This is commonly referred as "indirect reference" or "redirection."
	It's different from C, Go has not got the operation of pointer.
*/

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point i
	fmt.Println(*p) // get i by pointer p
	*p = 21         // set i by pointer p
	fmt.Println(i)  // output the value of i

	p = &j          // point j
	*p = *p / 37    // divide j by the pointer p
	fmt.Println(j)  // output the value of j
}
