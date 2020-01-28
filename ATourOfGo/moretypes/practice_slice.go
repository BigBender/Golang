/*
	Exercise: Slicing
	Implement Pic. 
	It should return a slice of length dy, where each element is a slice of length dx and element type uint8. 
	When you run this program, it will interpret each integer as a gray value (well, actually a blueness value) and display the image it corresponds to.

	The choice of image is up to you. Several interesting functions include (x + y) / 2, x * y, x ^ y, x * log (y), and x% (y + 1).

	Hint: you need to use a loop to assign each [] uint8 in [] [] uint8;
	use uint8 (intValue) to convert between types; you might use functions in the math package.
*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8,dy)    // outer slice
   for x := range a{
       b := make([]uint8,dx)   // inner slice
       for y := range b{
           b[y] = uint8(x%(y+1))
        }
        a[x] = b  
    }
    return a
}

func main() {
	pic.Show(Pic)
}
