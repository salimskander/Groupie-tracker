package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8000"

func GestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Gestion" {
		http.Error(w, "404 non reconnus", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Methode non supportée", http.StatusNotFound)
		return
	}

	fmt.Printf("Serveur opérationnel")

}

func main() {

	fileServer := http.FileServer(http.Dir("./Static/HTML/"))
	http.Handle("/", fileServer)

	http.HandleFunc("/Gestion", GestionHandler)

	fmt.Println("(http://localhost:8000/) - Serveur lancé sur le port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
