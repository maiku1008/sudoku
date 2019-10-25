package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	doku "github.com/sk1u/doku/internal"
)

func main() {
	gridstring := flag.String("s", "", "A grid in string form.")
	file := flag.String("f", "", "Text file with puzzles.")
	server := flag.Bool("server", false, "server toggle")
	flag.Parse()

	if *gridstring != "" {
		// TODO: Validate string
		Sudoku(*gridstring)
	}

	if *file != "" {
		lines, err := scanLines(*file)
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
	}

	if *server {
		mux := http.NewServeMux()
		mux.Handle("/newsudoku", doku.NewSudokuHandler())
		mux.Handle("/display", doku.NewDisplayHandler())
		mux.Handle("/solve", doku.NewSolveHandler())
		mux.Handle("/state", doku.NewStateHandler())

		fmt.Printf("Starting the server on port: 8080\n")
		err := http.ListenAndServe(":8080", mux)
		// TODO: do something useful with this error
		fmt.Println(err)
	}
}

// Sudoku has the main routine that can be called
func Sudoku(grid string) {
	s := doku.NewSudoku(grid)
	fmt.Println("Puzzle: ", grid)
	fmt.Println("")
	fmt.Println(s.Display())
	fmt.Println("Solving...")
	start := time.Now()
	s.Solve()
	elapsed := time.Since(start)
	fmt.Println("Solved: ", s.DisplayString())
	fmt.Println("")
	fmt.Println(s.Display())
	fmt.Println("Solved in: ", elapsed)
}

func scanLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
