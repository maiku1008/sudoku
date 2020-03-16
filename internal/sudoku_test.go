package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	easypuzzle       = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	mediumpuzzle     = "400000805030000000000700000020000060000080400000010000000603070500200000104000000"
	hardpuzzle       = "150300000070040200004072000008000000000900108010080790000003800000000000600007423"
	veryhardpuzzle   = "400000000000000000000000000000000000000000000000000000000000000000000000000000000"
	solvedpuzzle     = "417369825632158947958724316825437169791586432346912758289643571573291684164875293"
	impossiblepuzzle = "777777777777777777777777777777777777777777777777777777777777777777777777777777777"
)

var (
	c2units = [][]index{
		{"A1", "A2", "A3", "B1", "B2", "B3", "C1", "C2", "C3"},
		{"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"},
		{"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"},
	}
	c2peers = []index{
		"A1", "A2", "A3", "B1",
		"B2", "B3", "C1", "C3",
		"D2", "E2", "F2", "G2",
		"H2", "I2", "C4", "C5",
		"C6", "C7", "C8", "C9",
	}
)

func TestContains(t *testing.T) {
	assert := assert.New(t)

	d := value("123456789")
	assert.True(d.contains(value("3")))
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	d := value("123456789")
	d = d.remove("3")
	assert.Equal(d, value("12456789"))
}

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

func TestNewSudoku(t *testing.T) {
	assert := assert.New(t)
	s := NewSudoku(mediumpuzzle)

	// Testing s.squares
	assert.Equal(len(s.squares), 81)

	// Testing s.units
	for _, u := range c2units {
		assert.Contains(s.units["C2"], u)
	}

	// Testing s.peers
	for _, p := range c2peers {
		assert.Contains(s.peers["C2"], p)
	}
	assert.Equal(len(s.peers["C2"]), 20)

	// Testing s.grid
	assert.Equal(s.grid["A1"], value("4"))
	assert.Equal(s.grid["A2"], value("0"))
}

func TestIsSolved(t *testing.T) {
	assert := assert.New(t)

	assert.False(NewSudoku(veryhardpuzzle).isSolved())
	assert.True(NewSudoku(solvedpuzzle).isSolved())
}

func TestSolve(t *testing.T) {
	assert := assert.New(t)

	var puzzles = []struct {
		puzzleToSolve string
	}{
		{easypuzzle},
		{mediumpuzzle},
		{hardpuzzle},
		{veryhardpuzzle},
	}

	for _, p := range puzzles {
		s := NewSudoku(p.puzzleToSolve)
		err := s.Solve()
		assert.Nil(err)
		assert.True(s.isSolved())
	}

	s := NewSudoku(impossiblepuzzle)
	err := s.Solve()
	assert.Contains(err.Error(), "Contradiction: ")
}

func TestDisplayString(t *testing.T) {
	assert := assert.New(t)
	s := NewSudoku(mediumpuzzle)
	s.Solve()
	assert.Equal(s.DisplayString(), solvedpuzzle)
}
