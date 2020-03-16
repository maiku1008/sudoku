package sudoku

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	requestNewSudoku = "{\"grid\":\"400000805030000000000700000020000060000080400000010000000603070500200000104000000\"}"
	responseNewSudoku = "{\"hash\":\"TVAXF\",\"error\":\"None\"}\n"
	requestSolve = "{\"hash\":\"TVAXF\"}"
	responseSolve = "{\"error\":\"None\"}\n"
	requestState = "{\"hash\":\"TVAXF\"}"
	responseState = "{\"grid\":\"417369825632158947958724316825437169791586432346912758289643571573291684164875293\",\"solved\":true,\"error\":\"None\"}\n"
	requestWrongHash = "{\"hash\":\"EMASO\"}"
	responseWrongHash = "{\"grid\":\"\",\"solved\":false,\"error\":\"Sudoku not found\"}\n"
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

// A struct for testing the different endpoints, and their expected responses
var endpointTests = []struct {
	request string
	endpoint string
	response string
}{
	{requestNewSudoku, "/newsudoku", responseNewSudoku},
	{requestSolve, "/solve",responseSolve},
	{requestState, "/state",responseState},
	{requestWrongHash, "/state",responseWrongHash},
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
