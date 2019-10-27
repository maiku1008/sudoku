# sk1u/doku

## Intro

A sudoku solver written in Go, that can solve **any** puzzle;
It uses [Constraint Propagation](https://en.wikipedia.org/wiki/Constraint_satisfaction) and [Search](https://en.wikipedia.org/wiki/Search_algorithm) algorithms.

## Representation

We represent a puzzle with a 81 character long string.
It can use "." to represent an unknown value:

`..5...987.4..5...1..7......2...48....9.1.....6..2.....3..6..2.......9.7.......5..`

or a "0":

`600302000040000010000000000702600000000000054300000000080150000000040200000000700`

Some examples are included in the text files in this repository.

## Installation
Build with:

```
go build doku.go
```

Test with:
```
go test -v ./internal
```

## Use

We have several options to interface with the application.

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

Run with `-f filename.txt` for resolving the puzzles in each of the lines of the txt file
```
./doku -f puzzles_medium.txt
```

Run with `-server` flag to run a local webserver exposing API endpoints which wrap doku's main functions; useful for creating a full stack web application.

```
./doku -server
Starting the server on port: 8080
```

API endpoints are described below.

## API Endpoints
### /newsudoku
Calls the NewSudoku() method with `grid` as its argument.
Returns a `hash` that identifies the Sudoku object.
Use the same hash in subsequent calls to interface with your puzzle.

Request with `curl`:
```
curl -i \
-H "Accept: application/json" \
-X POST -d {\"grid\":\"600302000040000010000000000702600000000000054300000000080150000000040200000000700\"} \
http://localhost:8080/newsudoku
```

Example server response:
```
{
    "grid":"",
    "hash":"LAYHX",
    "solved":false,
    "error":""
}
```

### /display
Calls the DisplayString() method for the Sudoku object identified by `hash`.
Returns a `grid` string value representing the current Sudoku puzzle.

Request with `curl`:
```
curl -i \
-H "Accept: application/json" \
-X POST -d {\"hash\":\"LAYHX\"} \
http://localhost:8080/display
```

Example server response:
```
{
    "grid":"600302000040000010000000000702600000000000054300000000080150000000040200000000700",
    "hash":"",
    "solved":false,
    "error":""
}
```

### /solve
Calls the Solve() method for the Sudoku object identified by `hash`.
Doesn't return anything.

Request with `curl`:
```
curl -i \
-H "Accept: application/json" \
-X POST -d {\"hash\":\"LAYHX\"} \
http://localhost:8080/solve
```

Example server response:
```
{
    "grid":"",
    "hash":"",
    "solved":false,
    "error":""
}
```

### /state
Calls the isSolved() method for the Sudoku object identified by `hash`.
Returns a `solved` boolean field.

Request with `curl`:
```
curl -i \
-H "Accept: application/json" \
-X POST -d {\"hash\":\"LAYHX\"} \
http://localhost:8080/state
```

Example server response:
```
{
    "grid":"",
    "hash":"",
    "solved":true,
    "error":""
}
```

## Docker

Run the server app in a Docker container as such:
```
docker build -t doku .
docker run -d -p 8080:8080 doku
```

---

This project was inspired by Peter Norvig's [blog post](https://norvig.com/sudoku.html).