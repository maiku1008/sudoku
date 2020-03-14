package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	doku "github.com/micuffaro/doku/internal"
)

// Sudoku has the main routine that can be called
func Sudoku(grid string) {
	s := doku.NewSudoku(grid) // Create new sudoku object
	fmt.Println("Puzzle: ", grid)
	fmt.Println()
	fmt.Println(s.Display()) // Print a pretty display
	fmt.Println("Solving...")
	start := time.Now()
	s.Solve() // Actually solve
	elapsed := time.Since(start)
	fmt.Println("Solved: ", s.DisplayString())
	fmt.Println()
	fmt.Println(s.Display()) // Print a pretty display after solution
	fmt.Println("Solved in: ", elapsed)
}

// ReadFile reads from a file and outputs its contents
// to a slice, line by line
func ReadFile(path string) ([]string, error) {

	// Open the file found in path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Scan contents of file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	// Go through each token of the file and append it to slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func main() {

	// Parse arguments
	server := flag.Bool("server", false, "server toggle")
	gridString := flag.String("s", "", "A grid in string form.")
	file := flag.String("f", "", "Text file with puzzles.")
	flag.Parse()

	// Start a server which exposes the doku API
	if *server {
		mux := http.NewServeMux()
		mux.Handle("/newsudoku", doku.NewSudokuHandler())
		mux.Handle("/display", doku.NewDisplayHandler())
		mux.Handle("/solve", doku.NewSolveHandler())
		mux.Handle("/state", doku.NewStateHandler())

		fmt.Printf("Starting the server on port: 8080\n")
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			panic(err)
		}
	}

	// Solve the sudoku puzzle in gridString
	if *gridString != "" {
		Sudoku(*gridString)
	} else {
		fmt.Println("Invalid string for grid. Please specify one with -s")
		fmt.Println("Exiting")
		os.Exit(1)
	}

	// Solve the sudoku puzzles found in a file
	if *file != "" {
		lines, err := ReadFile(*file)
		if err != nil {
			panic(err)
		}

		start := time.Now()
		count := 0
		for _, line := range lines {
			Sudoku(line)
			fmt.Println("---")
			count++
		}
		elapsed := time.Since(start)
		fmt.Println("Solved", count, "Sudoku puzzles in", elapsed)
	} else {
		fmt.Println("Invalid file path. Please specify one with -f")
		fmt.Println("Exiting")
		os.Exit(1)
	}
}
