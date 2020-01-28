/*
	Let's do something fun with functions.

	Implement a fibonacci function that returns a function (closure) that returns a Fibonacci sequence (0, 1, 1, 2, 3, 5, ...).
*/

package main

import "fmt"

func fibonacci() func() int {
    back1, back2:= 0, 1  // Set two initial values in advance
    return func() int {        
        temp := back1    // Record the value of back1
        back1,back2 = back2,(back1 + back2)       
        return temp // return temp
    } 
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
	output
	0
	1
	1
	2
	3
	5
	8
	13
	21
	34
*/