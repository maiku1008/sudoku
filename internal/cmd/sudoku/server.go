package cmd

import (
	"fmt"
	sudoku "github.com/micuffaro/sudoku/internal"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a sudoku solving server",
	Long:  `Coming soon!`,
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.Handle("/newsudoku", sudoku.NewSudokuHandler(time.Now))
		mux.Handle("/solve", sudoku.NewSolveHandler())
		mux.Handle("/state", sudoku.NewStateHandler())

		fmt.Printf("Starting the server on port: 8080\n")
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			panic(err)
		}
	},
}
