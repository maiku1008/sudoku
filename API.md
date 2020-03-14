## API

Run with `-server` flag to run a local webserver exposing API endpoints which wrap sudoku's main functions; useful for creating a full stack web application.

```
./sudoku -server
Starting the server on port: 8080
```

## Endpoints
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
docker build -t sudoku .
docker run -d -p 8080:8080 sudoku
```

You can then access the described API via localhost:8080
