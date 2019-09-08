package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	helpUrl = "https://golang.org/doc/articles/wiki/"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path[1:]; path {
	case "":
		fmt.Fprintf(w, "Hello world!")
	case "help":
		fmt.Fprintf(w, helpUrl)
	default:
		fmt.Fprintf(w, "There is nothing at %s", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
