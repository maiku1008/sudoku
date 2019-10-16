package doku

import (
    "fmt"
    "net/http"
)

// TODO: remove this
const mediumpuzzle = "400000805030000000000700000020000060000080400000010000000603070500200000104000000"

// NewSudokuHandler initializes a sudokuHandler
func NewSudokuHandler() http.Handler {
    return sudokuHandler{}
}

// sudokuHandler ...
type sudokuHandler struct {
}

func (h sudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// NewDisplayHandler initializes a displayHandler
func NewDisplayHandler() http.Handler {
    return displayHandler{}
}

// displayHandler ...
type displayHandler struct {
}

func (h displayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := NewSudoku(mediumpuzzle)
    fmt.Println(s.Display())
}

// NewSolveHandler initializes a solveHandler
func NewSolveHandler() http.Handler {
    return solveHandler{}
}

// solveHandler ...
type solveHandler struct {
}

func (h solveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := NewSudoku(mediumpuzzle)
    s.Solve()
}

// NewStateHandler initializes a stateHandler
func NewStateHandler() http.Handler {
    return stateHandler{}
}

// stateHandler ...
type stateHandler struct {
}

func (h stateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := NewSudoku(mediumpuzzle)
    fmt.Println(s.isSolved())
}
