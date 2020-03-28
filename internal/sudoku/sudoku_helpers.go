package sudoku

import (
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

// isvalid is a helper function that determines if a value is ok to be parsed
// into our sudoku grid.
func isvalid(v string) bool {
	char := strings.Contains(string(digits), string(v))
	return char || v == "." || v == "0"
}
