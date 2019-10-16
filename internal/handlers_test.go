package doku

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Mux() *http.ServeMux {

    mux := http.NewServeMux()
    mux.Handle("/newsudoku", NewSudokuHandler())
    mux.Handle("/display", NewDisplayHandler())
    mux.Handle("/solve", NewSolveHandler())
    mux.Handle("/state", NewStateHandler())
    return mux
}

func TestEndPoints(t *testing.T) {
    assert := assert.New(t)

    puzzleJSON := "{\"grid\": \"400000805030000000000700000020000060000080400000010000000603070500200000104000000\"}"

    fmt.Println("POST /newsudoku", puzzleJSON)
    request, _ := http.NewRequest("POST", "/newsudoku", bytes.NewBuffer([]byte(puzzleJSON)))
    response := httptest.NewRecorder()
    Mux().ServeHTTP(response, request)
    assert.Equal(200, response.Code)

    fmt.Println("GET /display")
    request, _ = http.NewRequest("GET", "/display", nil)
    response = httptest.NewRecorder()
    Mux().ServeHTTP(response, request)
    assert.Equal(200, response.Code)

    fmt.Println("GET /state")
    request, _ = http.NewRequest("GET", "/state", nil)
    response = httptest.NewRecorder()
    Mux().ServeHTTP(response, request)
    assert.Equal(200, response.Code)

    fmt.Println("GET /solve")
    request, _ = http.NewRequest("GET", "/solve", nil)
    response = httptest.NewRecorder()
    Mux().ServeHTTP(response, request)
    assert.Equal(200, response.Code)
    // assert.Equal(response.Body.String(), "intro=\"1\"\n")

    fmt.Println("GET /display")
    request, _ = http.NewRequest("GET", "/display", nil)
    response = httptest.NewRecorder()
    Mux().ServeHTTP(response, request)
    assert.Equal(200, response.Code)

    fmt.Println("GET /state")
    request, _ = http.NewRequest("GET", "/state", nil)
    response = httptest.NewRecorder()
    Mux().ServeHTTP(response, request)
    assert.Equal(200, response.Code)
}
