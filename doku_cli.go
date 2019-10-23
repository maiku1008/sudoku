package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "time"

    "github.com/sk1u/doku/internal"
)

func main() {
    gridstring := flag.String("s", "", "A grid in string form.")
    file := flag.String("f", "", "Text file with puzzles.")
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
