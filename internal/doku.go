// Package doku solves sudoku puzzles
// through constraint propagation and search
package doku

// import (
//     "fmt"
//     "strings"
// )

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
    squares [81]index
    // A map of all units (row, column & box)
    // for each square
    units map[index][3][9]index
    // A map of all peers for each square
    peers map[index][20]index
}

// NewSudoku creates a new sudoku object given a grid string
// func NewSudoku(grid string) *Sudoku {
//     return &Sudoku
// }

func (s *Sudoku) isSolved() bool {
    return false
}

// Populates the sudoku's Units and Peers
func (s *Sudoku) populate() {
    // cols     = digits
    // squares  = cross(rows, cols)
    // unitlist = ([cross(rows, c) for c in cols] +
    //             [cross(r, cols) for r in rows] +
    //             [cross(rs, cs) for rs in ('ABC','DEF','GHI') for cs in ('123','456','789')])
    // units = dict((s, [u for u in unitlist if s in u])
    //              for s in squares)
    // peers = dict((s, set(sum(units[s],[]))-set([s]))
    //              for s in squares)
    // var (
    //     unitlist [][]index
    //     units    [3][9]index
    //     i        int
    // )
    // const (
    //     columns = "123456789"
    //     rows    = "ABCDEFGHI"
    // )
}
