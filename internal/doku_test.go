package doku

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCross(t *testing.T) {
    assert := assert.New(t)

    const (
        letters = "AB"
        numbers = "12"
    )

    var result = []index{"A1", "A2", "B1", "B2"}

    crossed := cross(letters, numbers)

    assert.Equal(len(crossed), 4)
    for i, v := range result {
        assert.Equal(crossed[i], v)
    }
}
