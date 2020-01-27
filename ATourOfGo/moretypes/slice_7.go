/*
	Sliced slice
	Slices can include any type, even including other slices.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a Tic-Tac-Toe Board (Classic Game)
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Two players take turns playing X and O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
// output
// X _ X
// O _ X
// _ _ O