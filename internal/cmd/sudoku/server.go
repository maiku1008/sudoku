package cmd

import (
	"fmt"
	"github.com/micuffaro/sudoku/internal/api"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

var (
	Port      string
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start a sudoku solving server",
		Long:  `Coming soon!`,
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			mux := http.NewServeMux()
			mux.Handle("/newsudoku", api.NewSudokuHandler(time.Now))
			mux.Handle("/solve", api.NewSolveHandler())
			mux.Handle("/state", api.NewStateHandler())

			fmt.Println("Starting the server on port:", port)
			port = fmt.Sprintf(":%v", port)
			err := http.ListenAndServe(port, mux)
			if err != nil {
				log.Fatalf("Error starting server: %v", err)
			}
		},
	}
)

func init() {
	serverCmd.Flags().StringVarP(&Port, "port", "p", "8080", "Port server should be listening to")
}
