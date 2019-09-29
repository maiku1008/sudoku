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

func (v value) remove(val value) value { return value(strings.Replace(string(v), string(val), "", -1)) }

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
// TODO: maybe display pretty or verbose according to flag
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
        // we print a dot for readability
        if value == "0" || len(value) > 1 {
            value = "."
        }
        grid += value + " "
    }
    grid += "\n"
    return grid
}

// Solve a sudoku by constraint propagation.
// The sudoku may not be entirely solved with only this solution.
func (s *Sudoku) Solve() error {
    if err := s.constraintPropagation(); err != nil {
        return err
    }

    if s.issolved() {
        return nil
    }
    // TODO: try with search
    return fmt.Errorf("Can't solve this")
}

func (s *Sudoku) constraintPropagation() error {
    tosolve := s.grid
    s.grid = make(grid)
    for _, square := range s.squares {
        s.grid[square] = digits
    }

    for i, value := range tosolve {
        // If the value is zero / unknown, we don't want to assign it
        if !strings.Contains(string(digits), string(value)) {
            continue
        }

        // fmt.Println(s.grid)
        if err := s.assign(value, i); err != nil {
            return err
        }
    }
    return nil
}

// assign doesn't really assign values,
// rather it eliminates all the other values
// except val from square i
func (s *Sudoku) assign(val value, i index) error {
    // Create a string of values to remove
    toremove := s.grid[i].remove(val)
    for _, rm := range toremove {
        // remove the values
        if err := s.eliminate(value(rm), i); err != nil {
            return err
        }
    }
    return nil
}

func (s *Sudoku) eliminate(val value, i index) error {
    // check if we already removed the value
    removed := strings.Contains(string(s.grid[i]), string(val))
    if !removed {
        return nil
    }

    // Remove the value val from the square i
    if len(s.grid[i]) > 1 {
        s.grid[i] = s.grid[i].remove(val)
    }

    // Check the length of values in the square i
    switch len(s.grid[i]) {
    case 0:
        return fmt.Errorf("removed last value from field %v", i)
    case 1:
        // If a square i is reduced to one value,
        // then eliminate it from its peers.
        if err := s.removeFromPeers(i); err != nil {
            return err
        }
    }

    for _, unit := range s.units[i] {
        // If a unit is reduced to only one place
        // for a value v, then put it there.
        found, place := s.singlePossibility(val, unit)
        if place == "" {
            return fmt.Errorf("no place for value %v is left", val)
        }
        if !found {
            continue
        }
        if err := s.assign(val, place); err != nil {
            return err
        }
    }

    return nil
}

// Removes the value of i from all its peers.
func (s *Sudoku) removeFromPeers(i index) error {
    for _, peer := range s.peers[i] {
        if err := s.eliminate(s.grid[i], peer); err != nil {
            return err
        }
    }
    return nil
}

// Returns true when the given value only has one possibility
// in the the square's units, and returns the index.
// If there is no possibility left, the index is empty string.
func (s *Sudoku) singlePossibility(val value, unit []index) (found bool, i index) {
    for _, u := range unit {
        ok := strings.Contains(string(s.grid[u]), string(val))
        if ok {
            // second possibility
            if found {
                return false, i
            }
            found = true
            i = u
        }
    }
    return found, i
}
