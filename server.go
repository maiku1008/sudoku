package main

import (
    "fmt"
    "net/http"

    "github.com/sk1u/doku/internal"
)

func main() {

    mux := http.NewServeMux()
    mux.Handle("/newsudoku", doku.NewSudokuHandler())
    mux.Handle("/display", doku.NewDisplayHandler())
    mux.Handle("/solve", doku.NewSolveHandler())
    mux.Handle("/state", doku.NewStateHandler())

    fmt.Printf("Starting the server on port: 8080\n")
    err := http.ListenAndServe(":8080", mux)
    // TODO: do something useful with this error
    fmt.Println(err)
}
