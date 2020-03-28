package api

import (
	"fmt"
	"github.com/micuffaro/sudoku/internal/sudoku"
	"net/http"
	"time"
)

const (
	NoError     = "None"
	Error       = "Sudoku not found"
	ContentType = "Content-Type"
	Application = "Application/json"
)

// sudokuStorage stores Sudoku objects
// TODO: Figure out better storage, like mongodb
var sudokuStorage = make(map[string]*sudoku.Sudoku)

// Request represents a request sent by the user
type Request struct {
	Grid string `json:"grid"`
	Hash string `json:"hash"`
}

// NewSudokuHandler initializes a newSudokuHandler
func NewSudokuHandler(timeFunc func() time.Time) http.Handler {
	return newSudokuHandler{timeFunc}
}

// JSON response for /newsudoku endpoint
type newSudokuResponse struct {
	Hash  string `json:"hash"`
	Error string `json:"error"`
}

type newSudokuHandler struct {
	// timeFunc is used to determine a time.Time object which
	// we will use to generate as a seed to generate unique hashes
	timeFunc func() time.Time
}

func (h newSudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := setRequest(r.Body)
	s := sudoku.NewSudoku(request.Grid)

	var response newSudokuResponse
	response.Hash = getHash(h.timeFunc())
	response.Error = NoError

	sudokuStorage[response.Hash] = s // Store the sudoku object
	setResponse(w, response)
}

// NewSolveHandler initializes a solveHandler
func NewSolveHandler() http.Handler { return solveHandler{} }

// JSON response for /solve endpoint
type solveResponse struct {
	Error string `json:"error"`
}

type solveHandler struct{}

func (h solveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := setRequest(r.Body)

	var response solveResponse
	if _, ok := sudokuStorage[request.Hash]; !ok {
		response.Error = Error
	} else {
		s := sudokuStorage[request.Hash]
		err := s.Solve()
		if err != nil {
			response.Error = fmt.Sprintf("%v", err)
		} else {
			response.Error = NoError
		}
	}
	setResponse(w, response)
}

// NewStateHandler initializes a stateHandler
func NewStateHandler() http.Handler { return stateHandler{} }

// JSON response for /state endpoint
type stateResponse struct {
	Grid   string `json:"grid"`
	Solved bool   `json:"solved"`
	Error  string `json:"error"`
}

type stateHandler struct{}

func (h stateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := setRequest(r.Body)

	var response stateResponse
	if _, ok := sudokuStorage[request.Hash]; !ok {
		response.Error = Error
	} else {
		s := sudokuStorage[request.Hash]
		response.Grid = s.DisplayString()
		response.Solved = s.IsSolved()
		response.Error = NoError
	}
	setResponse(w, response)
}
