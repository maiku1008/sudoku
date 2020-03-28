package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/micuffaro/sudoku/internal/api"
	"github.com/micuffaro/sudoku/internal/sudoku"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sync"
	"time"
)

var (
	String bool
	File   bool

	solveCmd = &cobra.Command{
		Use:   "solve",
		Short: "Solve a puzzle",
		Long:  `Coming soon!`,
		Args: func(cmd *cobra.Command, args []string) error {
			// Check that we actually get args
			if len(args) < 1 {
				return errors.New("Requires something to solve!")
			}

			// Ensure we don't get more than one arg
			if len(args) > 1 {
				return errors.New("Too many arguments")
			}

			// Get the flag value
			stringFlag, _ := cmd.Flags().GetBool("string")
			fileFlag, _ := cmd.Flags().GetBool("file")
			if stringFlag && fileFlag {
				return errors.New("Flags for string and file cannot both be true")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			object := args[0]
			stringFlag, _ := cmd.Flags().GetBool("string")
			fileFlag, _ := cmd.Flags().GetBool("file")

			if stringFlag && object != "" {
				if err := api.ValidateString(object); err != nil {
					fmt.Println(err)
					fmt.Println("Exiting...")
					os.Exit(1)
				}
				solveString(object)
				return
			}

			if fileFlag && object != "" {
				solveFile(object)
				return
			}
		},
	}
)

func init() {
	solveCmd.Flags().BoolVarP(&String, "string", "s", false, "A single string containing a sudoku puzzle")
	solveCmd.Flags().BoolVarP(&File, "file", "f", false, "Path to a file containing a list of puzzles on each line")
}

func solveString(grid string) {
	s := sudoku.NewSudoku(grid) // Create new sudoku object
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

// Solve the sudoku puzzles found in a file
func solveFile(path string) {
	lines, err := readFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var wg sync.WaitGroup

	start := time.Now()
	count := 0
	for _, line := range lines {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			solveString(l)
			fmt.Println("---")
			count++
		}(line)
	}
	wg.Wait() // Wait for goroutines to finish
	elapsed := time.Since(start)
	fmt.Println("Solved", count, "Sudoku puzzles in", elapsed)
}

// ReadFile reads from a file and outputs its contents
// to a slice, line by line
func readFile(path string) ([]string, error) {

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
