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
    // Peers are all the squares that share a unit
    peers map[index][]index
}

// NewSudoku creates a new sudoku object given a grid string
func NewSudoku(grid string) *Sudoku {
    s := Sudoku{
        squares: cross(index("ABCDEFGHI"), index("123456789")),
    }
    s.populate()
    s.parse(grid)
    return &s
}

// Parse the sudoku from a string. The string
// should have either 0s or . for empty fields.
func (s *Sudoku) parse(grid string) {
    s.grid = make(map[index]value)
    i := 0
    for _, v := range grid {
        val := value(v)
        if ok := isvalid(string(v)); !ok {
            continue
        }

        // if val == value("0") || val == value(".") {
        //     s.grid[s.squares[i]] = value("123456789")
        // } else {
        //     s.grid[s.squares[i]] = val
        // }
        // TODO: This doesnt assign the whole range of values
        s.grid[s.squares[i]] = val
        i++
    }
}

// A helper function that returns true if the Sudoku is solved
func (s *Sudoku) isSolved() bool {
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
    width := 0
    for _, square := range s.squares {
        if width < len(s.grid[square]) {
            width = len(s.grid[square])
        }
    }
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
        // Print a dot instead of 0 for readability
        if value == "0" {
            value = "."
        }

        // Center the value if needed
        if len(value) < width {
            value = fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(value))/2, value))
        }
        grid += value + " "
    }
    grid += "\n"
    return grid
}

// Solve a sudoku by constraint propagation and search
func (s *Sudoku) Solve() error {
    if err := s.constraintPropagation(); err != nil {
        return err
    }

    if s.isSolved() {
        return nil
    }
    return s.search()
}

// Attempt to solve the sudoku with search (guessing)
func (s *Sudoku) search() error {
    // Choose the unfilled square with the fewest possibilities
    min := s.minimumValues()
    // Try through the possible values
    for _, val := range s.grid[min] {
        sc, err := s.try(value(val), min)
        if err != nil {
            continue
        }
        // If we got here, that means we have a solution.
        *s = *sc
        return nil
    }
    return fmt.Errorf("All possibilties lead nowhere")
}

// minimumValues gets the square that has the lowest
// number of possible values
func (s *Sudoku) minimumValues() index {
    var minField index
    minPoss := 10 // will never have more than 9 possibilities
    for _, square := range s.squares {
        numPoss := len(s.grid[square])
        // we already know what's in there
        if numPoss == 1 {
            continue
        }
        if numPoss < minPoss {
            minField = square
            minPoss = numPoss
        }
    }
    return minField
}

// Try to set a value at the given square, returning a copy of
// the sudoku with the tried value filled in, or an error if
// a contradiction has been detected, making this move invalid.
func (s *Sudoku) try(val value, i index) (*Sudoku, error) {
    // create a copy
    sc := s.copy()
    if err := sc.assign(val, i); err != nil {
        return nil, err
    }
    if !sc.isSolved() {
        if err := sc.search(); err != nil {
            return nil, err
        }
    }
    return sc, nil
}

// Creates a copy of the Sudoku, including a deep-copy
// of the main sudoku grid. We use this when doing guesswork
// that could be wrong.
func (s Sudoku) copy() *Sudoku {
    newGrid := make(grid)
    for k, v := range s.grid {
        newGrid[k] = v
    }
    s.grid = newGrid
    return &s
}

// Attempt to solve the sudoku through constraint propagation.
// Harder puzzles may not be solved by this method alone.
func (s *Sudoku) constraintPropagation() error {
    tosolve := s.grid
    s.grid = make(grid)
    for _, square := range s.squares {
        s.grid[square] = digits
    }

    for i, value := range tosolve {
        ok := strings.Contains(string(digits), string(value))
        if !ok {
            continue
        }

        if err := s.assign(value, i); err != nil {
            return err
        }
    }
    return nil
}

// Assign doesn't really assign values,
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

// Eliminate a value from an index, and propagate.
// This is the main block of the constraint propagation method
// of solving.
func (s *Sudoku) eliminate(val value, i index) error {

    // check if we already removed the value
    removed := strings.Contains(string(s.grid[i]), string(val))
    if !removed {
        return nil
    }

    // Remove the value val from the square i
    s.grid[i] = s.grid[i].remove(val)

    // Check the length of values in the square i
    switch len(s.grid[i]) {
    case 0:
        return fmt.Errorf("Contradiction: removed last value from field %v", i)
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
            return fmt.Errorf("Contradiction: no place for value %v is left", val)
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

// Removes the value of square i from all its peers.
func (s *Sudoku) removeFromPeers(i index) error {
    for _, peer := range s.peers[i] {
        if err := s.eliminate(s.grid[i], peer); err != nil {
            return err
        }
    }
    return nil
}

// Returns true when the given value only has one possibility
// in the the square's units, and returns its index.
// If there is no possibility left, the index is an mpty string.
func (s *Sudoku) singlePossibility(val value, unit []index) (found bool, i index) {
    for _, u := range unit {
        ok := strings.Contains(string(s.grid[u]), string(val))
        if ok {
            if found {
                return false, i
            }
            found = true
            i = u
        }
    }
    return found, i
}
