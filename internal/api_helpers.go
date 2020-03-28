package sudoku

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// setRequest takes the io.ReadCloser object and
// unmarshals it into a Request struct
func setRequest(b io.ReadCloser) Request {

	// Read the body
	body, err := ioutil.ReadAll(b)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	// Unmarshal the body into the request struct
	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Unmarshalling error: %v", err)
	}

	return request
}

// setResponse encodes the response to json and writes it to w
func setResponse(w http.ResponseWriter, response interface{}) {

	// Set header
	w.Header().Set(ContentType, Application)
	w.WriteHeader(http.StatusCreated)

	// Encode the JSON
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Encoding error: %v", err)
	}
}

// getHash generates a random 5 digit hash according time
func getHash(t time.Time) string {
	rand.Seed(t.UnixNano())
	return randomString(5)
}
