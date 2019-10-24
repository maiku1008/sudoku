# sk1u/doku

## Intro

A sudoku solver written in Go, that can solve *any* puzzle;
It uses [Constraint Propagation](https://en.wikipedia.org/wiki/Constraint_satisfaction) and [Search](https://en.wikipedia.org/wiki/Search_algorithm) algorithms.

## Representation

We represent a puzzle with a 81 character long string.

examples:
`..5...987.4..5...1..7......2...48....9.1.....6..2.....3..6..2.......9.7.......5..`
`600302000040000010000000000702600000000000054300000000080150000000040200000000700`

## Installation

```
go build doku.go
```

## Use

We interface with this application with a simple cli app for running locally.

Run with `-s` flag to solve a single puzzle.

```
./doku -s "..5...987.4..5...1..7......2...48....9.1.....6..2.....3..6..2.......9.7.......5.."
```
Output:
```
Puzzle:  ..5...987.4..5...1..7......2...48....9.1.....6..2.....3..6..2.......9.7.......5..

 . . 5 | . . . | 9 8 7
 . 4 . | . 5 . | . . 1
 . . 7 | . . . | . . .
-------+-------+-------
 2 . . | . 4 8 | . . .
 . 9 . | 1 . . | . . .
 6 . . | 2 . . | . . .
-------+-------+-------
 3 . . | 6 . . | 2 . .
 . . . | . . 9 | . 7 .
 . . . | . . . | 5 . .

Solving...
Solved:  135426987846957321927381465213748659598163742674295813351674298482539176769812534

 1 3 5 | 4 2 6 | 9 8 7
 8 4 6 | 9 5 7 | 3 2 1
 9 2 7 | 3 8 1 | 4 6 5
-------+-------+-------
 2 1 3 | 7 4 8 | 6 5 9
 5 9 8 | 1 6 3 | 7 4 2
 6 7 4 | 2 9 5 | 8 1 3
-------+-------+-------
 3 5 1 | 6 7 4 | 2 9 8
 4 8 2 | 5 3 9 | 1 7 6
 7 6 9 | 8 1 2 | 5 3 4

Solved in:  11.149529ms
```

Run with `-server` flag to run a server option where we can wrap its main functions in API endpoints, useful for creating a full stack web application.


## TODO:
- [ ] Complete Readme
- [ ] Add Go CI
- [x] Write simple CLI app
- [x] Validate if hash exists in server app
- [x] Write a DisplayString() that will simply output the solved sudoku in string form
- [x] Write a test and a handler for the above DisplayString()
- [x] Consider splitting up the session struct into a request and response struct instead
- [ ] Dockerize server app
- [ ] Use different storage other than a map

---

This project has been inspired by [this](https://norvig.com/sudoku.html) blog post.