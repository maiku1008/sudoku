package sudoku

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	NoError     = "None"
	ContentType = "Content-Type"
	Application = "Application/json"
)

// sudokuStorage stores Sudoku objects
// TODO: Figure out better storage, like mongodb
var sudokuStorage = make(map[string]*Sudoku)

// Request represents a request sent by the user
type Request struct {
	Grid   string `json:"grid"`
	Hash   string `json:"hash"`
}

// dumpJSON takes the io.ReadCloser object and
// unmarshals it into a Request struct
func dumpRequest(b io.ReadCloser) Request {

	// Read the body
	body, err := ioutil.ReadAll(b)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	// Unmarshal the body into the request struct
	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Fatalf("Unmarshalling error: %v", err)
	}

	return request
}

// NewSudokuHandler initializes a newSudokuHandler
func NewSudokuHandler() http.Handler { return newSudokuHandler{} }

// JSON response for /newsudoku endpoint
type newSudokuResponse struct {
	Hash   string `json:"hash"`
	Error  string `json:"error"`
}

type newSudokuHandler struct{}

func (h newSudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := dumpRequest(r.Body)
	s := NewSudoku(request.Grid)

	var response newSudokuResponse
	rand.Seed(time.Now().UnixNano()) // Generate a random seed according to current time
	response.Hash = randomString(5)
	response.Error = NoError

	sudokuStorage[response.Hash] = s // Store the sudoku object

	w.Header().Set(ContentType, Application)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding error: %v", err)
	}
}

// NewDisplayHandler initializes a displayHandler
func NewDisplayHandler() http.Handler { return displayHandler{} }

// JSON response for display endpoint
type displayResponse struct {
	Grid   string `json:"grid"`
	Error  string `json:"error"`
}

type displayHandler struct{}

func (h displayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := dumpRequest(r.Body)

	var response displayResponse
	if _, ok := sudokuStorage[request.Hash]; !ok {
		response.Error = "Sudoku not found"
	} else {
		s := sudokuStorage[request.Hash]
		response.Grid = s.DisplayString()
		response.Error = NoError
	}

	w.Header().Set(ContentType, Application)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding error: %v", err)
	}
}

// NewSolveHandler initializes a solveHandler
func NewSolveHandler() http.Handler { return solveHandler{} }

// JSON response for /solve endpoint
type solveResponse struct {
	Error  string `json:"error"`
}

type solveHandler struct{}

func (h solveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := dumpRequest(r.Body)

	var response solveResponse
	if _, ok := sudokuStorage[request.Hash]; !ok {
		response.Error = "Sudoku not found"
	} else {
		s := sudokuStorage[request.Hash]
		err := s.Solve()
		if err != nil {
			response.Error = fmt.Sprintf("%v", err)
		} else {
			response.Error = NoError
		}
	}

	w.Header().Set(ContentType, Application)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding error: %v", err)
	}
}

// NewStateHandler initializes a stateHandler
func NewStateHandler() http.Handler { return stateHandler{} }

// JSON response for /state endpoint
type stateResponse struct {
	Solved bool   `json:"solved"`
	Error  string `json:"error"`
}

type stateHandler struct{}

func (h stateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request := dumpRequest(r.Body)

	var response stateResponse
	if _, ok := sudokuStorage[request.Hash]; !ok {
		response.Error = "Sudoku not found"
	} else {
		s := sudokuStorage[request.Hash]
		response.Solved = s.isSolved()
		response.Error = NoError
	}

	w.Header().Set(ContentType, Application)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding error: %v", err)
	}
}
