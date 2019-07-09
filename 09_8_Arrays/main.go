// Slices of lsices

// slices can contain any type including other slices

package main

import (
	"fmt"
	"strings"
)

func main() {
	//create a tic-tac-toe board

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//The playes take turn

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "x"
	board[1][0] = "O"
	board[0][2] = "x"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
