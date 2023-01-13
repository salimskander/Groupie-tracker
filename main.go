package main

import (
	"fmt"
	controler "groupie/controler"
	"log"
	"net/http"
)

const port = ":8080"

func GestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Gestion" {
		http.Error(w, "404 non reconnus", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Methode non supportée", http.StatusNotFound)
		return
	}
}

func main() {

	fs := http.FileServer(http.Dir("Static/HTML/"))
	http.Handle("/Static/CSS/", http.StripPrefix("/Static/CSS/", fs))
	http.HandleFunc("/Gestion", GestionHandler)
	http.HandleFunc("/", controler.HomePage)
	fmt.Println("(http://localhost:8080/) - Serveur lancé sur le port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
		fmt.Println("Fatal error serveur ne se lance pas")
	}

}
