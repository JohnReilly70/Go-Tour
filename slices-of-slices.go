package main

import (
	"fmt"
	"strings"
)

func main() {
	//creating tic tac board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//player turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Println(i)
		fmt.Printf("%b BINARY %v VALUE %T TYPE\n", i, i, i)
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}
