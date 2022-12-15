package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported.", http.StatusNotFound)
		return
	}
	fmt.Printf("Everything is good\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
