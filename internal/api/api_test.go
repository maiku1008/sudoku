package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Mux() *http.ServeMux {
	// We use this to generate always the same hash when adding a new grid
	fixedTime := func() time.Time {
		return time.Date(1983, 9, 1, 20, 0, 0, 0, time.UTC)
	}

	mux := http.NewServeMux()
	mux.Handle("/newsudoku", NewSudokuHandler(fixedTime))
	mux.Handle("/solve", NewSolveHandler())
	mux.Handle("/state", NewStateHandler())
	return mux
}

const (
	// Valid request to /newsudoku
	requestNewSudoku  = "{\"grid\":\"400000805030000000000700000020000060000080400000010000000603070500200000104000000\"}"
	responseNewSudoku = "{\"hash\":\"TVAXF\",\"error\":\"None\"}\n"

	// Invalid requests to /newsudoku : Invalid grid length
	requestNewSudokuInvalidLength  = "{\"grid\":\"\"}"
	responseNewSudokuInvalidLength = "{\"hash\":\"\",\"error\":\"Wrong length for sudoku grid\"}\n"

	// Invalid requests to /newsudoku : Invalid characters in grid
	requestNewSudokuInvalidCharacters  = "{\"grid\":\"ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234567890123456789012345678901234567890ABCDE\"}"
	responseNewSudokuInvalidCharacters = "{\"hash\":\"\",\"error\":\"Invalid characters in sudoku grid\"}\n"

	// Valid request to /solve
	requestSolve  = "{\"hash\":\"TVAXF\"}"
	responseSolve = "{\"error\":\"None\"}\n"

	// Invalid request to /solve : Provided hash doesnt exist in storage
	requestSolveInvalidHash  = "{\"hash\":\"EMASO\"}"
	responseSolveInvalidHash = "{\"error\":\"Sudoku not found\"}\n"

	// Valid request to /state
	requestState  = "{\"hash\":\"TVAXF\"}"
	responseState = "{\"grid\":\"417369825632158947958724316825437169791586432346912758289643571573291684164875293\",\"solved\":true,\"error\":\"None\"}\n"

	// Invalid request to /state : Provided hash doesnt exist in storage
	requestStateInvalidHash  = "{\"hash\":\"EMASO\"}"
	responseStateInvalidHash = "{\"grid\":\"\",\"solved\":false,\"error\":\"Sudoku not found\"}\n"
)

// A struct for testing the different endpoints, and their expected responses
var endpointTests = []struct {
	request  string
	endpoint string
	response string
}{
	{requestNewSudoku, "/newsudoku", responseNewSudoku},
	{requestNewSudokuInvalidLength, "/newsudoku", responseNewSudokuInvalidLength},
	{requestNewSudokuInvalidCharacters, "/newsudoku", responseNewSudokuInvalidCharacters},
	{requestSolve, "/solve", responseSolve},
	{requestSolveInvalidHash, "/solve", responseSolveInvalidHash},
	{requestState, "/state", responseState},
	{requestStateInvalidHash, "/state", responseStateInvalidHash},
}

func TestEndPoints(t *testing.T) {
	assert := assert.New(t)

	for _, et := range endpointTests {
		r := bytes.NewBuffer([]byte(et.request))
		request, _ := http.NewRequest("POST", et.endpoint, r)
		response := httptest.NewRecorder()
		Mux().ServeHTTP(response, request)

		assert.Equal(et.response, response.Body.String())
	}
}
