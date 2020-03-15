package sudoku

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Mux() *http.ServeMux {

	mux := http.NewServeMux()
	mux.Handle("/newsudoku", NewSudokuHandler())
	mux.Handle("/solve", NewSolveHandler())
	mux.Handle("/state", NewStateHandler())
	return mux
}

// A struct for testing the different endpoints
var endpointTests = []struct {
	method   string
	endpoint string
	respcode int
}{
	{"POST", "/state", 201},
	{"POST", "/solve", 201},
	{"POST", "/state", 201},
}

func TestEndPoints(t *testing.T) {
	assert := assert.New(t)

	puzzle1 := "{\"grid\": \"400000805030000000000700000020000060000080400000010000000603070500200000104000000\"}"
	b := bytes.NewBuffer([]byte(puzzle1))

	fmt.Println("POST /newsudoku", puzzle1)
	request, _ := http.NewRequest("POST", "/newsudoku", b)
	response := httptest.NewRecorder()
	Mux().ServeHTTP(response, request)
	assert.Equal(201, response.Code)

	requestBody := response.Body.String()

	for _, et := range endpointTests {
		JSON := bytes.NewBuffer([]byte(requestBody))

		fmt.Println("Request:\n", et.method, et.endpoint, JSON)
		request, _ = http.NewRequest(et.method, et.endpoint, JSON)
		response = httptest.NewRecorder()
		Mux().ServeHTTP(response, request)
		assert.Equal(et.respcode, response.Code)
		fmt.Println("Response:\n", response.Body)
	}
}

func TestEndpointInvalidHash(t *testing.T) {
	assert := assert.New(t)

	jsonString := "{\"grid\":\"400000805030000000000700000020000060000080400000010000000603070500200000104000000\",\"hash\":\"EMASO\"}"
	JSON := bytes.NewBuffer([]byte(jsonString))
	request, _ := http.NewRequest("POST", "/state", JSON)
	response := httptest.NewRecorder()
	Mux().ServeHTTP(response, request)
	assert.Contains(response.Body.String(), "Sudoku not found")
}

// TODO:Test each endpoint with request/response