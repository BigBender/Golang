/*
	Array
	Type [n]T representing n values of type T

	Expression var a [10]int
	The variable a is declared as an array of 10 integers.

	Array length cannot be changed
*/

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
