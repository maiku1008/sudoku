package main

import (
	cmd "github.com/micuffaro/sudoku/internal/cmd/sudoku"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
