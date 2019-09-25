package doku

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestFind(t *testing.T) {
    assert := assert.New(t)

    var i int
    var found bool

    var example = []index{"0", "1", "7", "2", "3", "4"}
    var findTests = []struct {
        elementToFind index
        expectedIndex int
        expectedFound bool
    }{
        {"7", 2, true},
        {"14", -1, false},
    }

    for _, ft := range findTests {
        i, found = find(example, ft.elementToFind)
        assert.Equal(i, ft.expectedIndex)
        assert.Equal(found, ft.expectedFound)
    }
}

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

func TestSudoku(t *testing.T) {
    assert := assert.New(t)
    s := NewSudoku()

    // Testing s.squares
    assert.Equal(len(s.squares), 81)
    // fmt.Println("squares: ", s.squares)

    // Testing s.units
    c2units := [][]index{
        {"A1", "A2", "A3", "B1", "B2", "B3", "C1", "C2", "C3"},
        {"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"},
        {"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"},
    }
    for _, u := range c2units {
        assert.Contains(s.units["C2"], u)
    }
    // fmt.Println(s.units["A2"])

    // Testing s.peers
    c2peers := []index{
        "A1", "A2", "A3", "B1",
        "B2", "B3", "C1", "C3",
        "D2", "E2", "F2", "G2",
        "H2", "I2", "C4", "C5",
        "C6", "C7", "C8", "C9",
    }
    for _, p := range c2peers {
        assert.Contains(s.peers["C2"], p)
    }
    assert.Equal(len(s.peers["C2"]), 20)
    // fmt.Println(s.peers["A2"])
}
