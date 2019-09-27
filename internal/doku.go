// Package doku solves sudoku puzzles
// through constraint propagation and search
package doku

import (
    "fmt"
    "strings"
)

const (
    digits = value("123456789")
)

// The index of our map, ex. "A1"
type index string

// All possible values for a square, ex. "123456789"
type value string

// Our full sudoku grid
type grid map[index]value

// Sudoku is our sudoku object
type Sudoku struct {
    // The actual grid
    grid
    // All squares in the grid
    squares []index
    // A map of all units (row, column & box)
    // for each square
    units map[index][][]index
    // A map of all peers for each square
    peers map[index][]index
}

// NewSudoku creates a new sudoku object given a grid string
func NewSudoku(customgrid string) *Sudoku {
    s := Sudoku{
        squares: cross(index("ABCDEFGHI"), index("123456789")),
    }
    s.populate()
    s.parse(customgrid)
    return &s
}

func isvalid(v string) bool {
    char := strings.Contains(string(digits), string(v))
    return char || v == "." || v == "0"
}

// Parse the sudoku from a string. The string
// should have either 0s or '.' for empty fields,
// everything else gets ignored
func (s *Sudoku) parse(customgrid string) {
    s.grid = make(map[index]value)
    i := 0
    for _, v := range customgrid {
        val := value(v)
        if ok := isvalid(string(v)); !ok {
            continue
        }

        if val == value("0") || val == value(".") {
            s.grid[s.squares[i]] = value("123456789")
        } else {
            s.grid[s.squares[i]] = val
        }
        i++
    }
}

// A helper function that returns true if the Sudoku is solved
func (s *Sudoku) issolved() bool {
    for _, square := range s.grid {
        if len(square) > 1 {
            return false
        }
    }
    return true
}

// Populates the sudoku's Units and Peers
func (s *Sudoku) populate() {
    const (
        columns = "123456789"
        rows    = "ABCDEFGHI"
    )

    // Build our Unit List
    unitlist := make([][]index, 27)
    i := 0
    for _, r := range []index{"ABC", "DEF", "GHI"} {
        for _, c := range []index{"123", "456", "789"} {
            unitlist = append(
                unitlist,
                cross(rows, index(columns[i])),
                cross(index(rows[i]), columns),
                cross(r, c),
            )
            i++
        }
    }

    // Populate our units and peers
    s.units = make(map[index][][]index, 81)
    s.peers = make(map[index][]index, 81)

    for _, square := range s.squares {
        for _, unit := range unitlist {
            _, ok := find(unit, square)
            if ok {
                s.units[square] = append(s.units[square], unit)
            }
        }
        for _, unit := range s.units[square] {
            for _, u := range unit {
                _, ok := find(s.peers[square], u)
                if !ok && square != u {
                    s.peers[square] = append(s.peers[square], u)
                }
            }
        }
    }
}

// Display pretty-prints the sudoku.
func (s *Sudoku) Display() string {
    width := 1
    line := strings.Repeat("-", ((width+1)*3)+1)
    line = fmt.Sprintf("%v+%v+%v", line, line, line)
    var grid string
    for i, square := range s.squares {
        switch {
        case i == 0:
            grid += " "
        case (i % 27) == 0:
            grid += "\n" + line + "\n "
        case (i % 9) == 0:
            grid += "\n "
        case (i % 3) == 0:
            grid += "| "
        }
        value := string(s.grid[square])
        // If each square has more than one possible value,
        // we print a dot
        if value == "0" || value == "123456789" {
            value = "."
        }
        grid += value + " "
    }
    grid += "\n"
    return grid
}
