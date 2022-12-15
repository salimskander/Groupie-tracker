package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./HTML/index.html")
	if err != nil {
		panic(err)
	}

	tmp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("server started at http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
