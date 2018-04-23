package main

import (
	"net/http"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("Please specify file directory.")
	}
	http.Handle("/", http.FileServer(http.Dir(args[1])))
	http.ListenAndServe(":9090", nil)
}
