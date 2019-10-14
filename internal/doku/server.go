package doku

// We want this program to create the following http handlers:

// NewSudoku -> takes a grid string and initializes our sudoku
// Display -> returns a string representing our sudoku
// Solve -> Solves our Sudoku
// IsSolved -> returns true if the sudoku is solved

import (
    "net/http"
)

// NewSudokuHandler ...
type NewSudokuHandler struct {
}

func (h NewSudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// DisplayHandler ...
type DisplayHandler struct {
}

func (h DisplayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := NewSudoku(mediumpuzzle)
    s.Display()
}

// SolveHandler ...
type SolveHandler struct {
}

func (h SolveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := NewSudoku(mediumpuzzle)
    s.Solve()
}

// StateHandler ...
type StateHandler struct {
}

func (h StateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := NewSudoku(mediumpuzzle)
    s.Solve()
}
