package main

import (
    "fmt"
    "net/http"

    "github.com/sk1u/doku/internal"
)

func main() {

    // fileJSON := flag.String("json", "web/gopher.json", "JSON file")
    // fileHTML := flag.String("template", "web/index.html", "HTML Template file")
    // port := flag.String("port", "9999", "Webserver port")
    // flag.Parse()

    // storyJSON := InitStoryJSON(*fileJSON)
    // template := InitTemplate(*fileHTML)

    mux := http.NewServeMux()
    mux.Handle("/newsudoku", doku.NewSudokuHandler())
    mux.Handle("/display", doku.DisplayHandler())
    mux.Handle("/solve", doku.SolveHandler())
    mux.Handle("/state", doku.StateHandler())

    fmt.Printf("Starting the server on port: 8080\n")
    err := http.ListenAndServe(":8080", mux)
    // TODO: do something useful with this error
    fmt.Println(err)
}
