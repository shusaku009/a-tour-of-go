package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "x"
	board[2][2] = "O"
	board[1][2] = "x"
	board[1][0] = "x"
	board[0][2] = "x"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
