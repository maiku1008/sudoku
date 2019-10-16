package doku

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const puzzle = "400000805030000000000700000020000060000080400000010000000603070500200000104000000"

// TODO: Figure out better storage, like mongodb
var sudokuStorage = make(map[string]*Sudoku)

// UserGrid is used to map data for grid
type UserGrid struct {
    Grid string `json:"grid"`
}

// NewSudokuHandler initializes a sudokuHandler
func NewSudokuHandler() http.Handler {
    return sudokuHandler{}
}

// sudokuHandler ...
type sudokuHandler struct {
    grid string
}

func (h sudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    body, _ := ioutil.ReadAll(r.Body)

    var usergrid UserGrid
    err := json.Unmarshal(body, &usergrid)
    if err != nil {
        fmt.Println(err)
    }

    h.grid = usergrid.Grid
    s := NewSudoku(h.grid)

    sudokuStorage[h.grid] = s
}

// NewDisplayHandler initializes a displayHandler
func NewDisplayHandler() http.Handler {
    return displayHandler{}
}

// displayHandler ...
type displayHandler struct {
}

func (h displayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s := sudokuStorage[puzzle]
    fmt.Fprintf(w, s.Display())
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
    s := sudokuStorage[puzzle]
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
    s := sudokuStorage[puzzle]
    fmt.Fprintf(w, "%t", s.isSolved())
    fmt.Println(s.isSolved())
}
