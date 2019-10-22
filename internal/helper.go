package doku

import (
    "math/rand"
    "strings"
)

// cross populates a slice of indexes with
// the cross between A and B
func cross(A, B index) []index {
    result := make([]index, len(A)*len(B))
    i := 0
    for _, a := range A {
        for _, b := range B {
            result[i] = index(a) + index(b)
            i++
        }
    }
    return result
}

// find is a helper function that takes a slice and looks for an element val in it.
// If found it will return its key, otherwise it will return -1 and a bool of false.
func find(slice []index, val index) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

// isvalid is a helper function that determins if a value is ok to be parsed
// into our sudoku grid.
func isvalid(v string) bool {
    char := strings.Contains(string(digits), string(v))
    return char || v == "." || v == "0"
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}
