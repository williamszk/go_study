package main

import (
	"dengo/pkg/node_types"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// server()
	cli_arg := cli_args()
	runChosenNodeType(cli_arg)
}

func runChosenNodeType(cli_arg string) {
	nodeType := node_types.FromString(cli_arg)
	switch nodeType {
	case node_types.Worker:
		fmt.Println("Selected Worker")
	case node_types.Manager:
		fmt.Println("Selected Manager")
	default:
		panic("Sorry")
	}
}

func runWorker() {

	// Define a handler function for the root endpoint ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func runManager() {

}

func cli_args() string {
	// Parse command-line flags
	flag.Parse()

	// Access non-flag arguments by position
	args := flag.Args()

	// Check the number of arguments
	if len(args) != 1 {
		fmt.Println("Usage: program <arg1>")
		os.Exit(1)
	}
	// Access arguments by position
	arg1 := args[0]
	// arg2 := args[1]

	return arg1
}
