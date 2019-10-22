package doku

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "math/rand"
    "net/http"
    "time"
)

// sudokuStorage stores Sudoku objects
// TODO: Figure out better storage, like mongodb
var sudokuStorage = make(map[string]*Sudoku)

// jsonStruct is used to map the json data for the session.
type jsonStruct struct {
    Grid   string `json:"grid"`
    Hash   string `json:"hash"`
    Solved bool   `json:"solved"`
    Error  string `json:"error"`
}

func dumpJSON(b io.ReadCloser) jsonStruct {

    body, _ := ioutil.ReadAll(b)

    var jsonStruct jsonStruct
    err := json.Unmarshal(body, &jsonStruct)
    if err != nil {
        fmt.Println(err)
    }

    return jsonStruct
}

// NewSudokuHandler initializes a sudokuHandler
func NewSudokuHandler() http.Handler { return sudokuHandler{} }

type sudokuHandler struct{}

func (h sudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    request := dumpJSON(r.Body)

    s := NewSudoku(request.Grid)

    var response jsonStruct
    rand.Seed(time.Now().UnixNano())
    response.Hash = randomString(5)
    sudokuStorage[response.Hash] = s

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// NewDisplayHandler initializes a displayHandler
func NewDisplayHandler() http.Handler { return displayHandler{} }

type displayHandler struct{}

func (h displayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    request := dumpJSON(r.Body)

    var response jsonStruct
    if _, ok := sudokuStorage[request.Hash]; !ok {
        response.Error = "Sudoku not found"
    } else {
        s := sudokuStorage[request.Hash]
        response.Grid = s.DisplayString()
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// NewSolveHandler initializes a solveHandler
func NewSolveHandler() http.Handler { return solveHandler{} }

type solveHandler struct{}

func (h solveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    request := dumpJSON(r.Body)

    var response jsonStruct
    if _, ok := sudokuStorage[request.Hash]; !ok {
        response.Error = "Sudoku not found"
    } else {
        s := sudokuStorage[request.Hash]
        s.Solve()
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// NewStateHandler initializes a stateHandler
func NewStateHandler() http.Handler { return stateHandler{} }

type stateHandler struct{}

func (h stateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    request := dumpJSON(r.Body)

    var response jsonStruct
    if _, ok := sudokuStorage[request.Hash]; !ok {
        response.Error = "Sudoku not found"
    } else {
        s := sudokuStorage[request.Hash]
        response.Solved = s.isSolved()
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}
