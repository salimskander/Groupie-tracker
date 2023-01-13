package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Accueil struct {
	Id_groupe     int    `json:"id"`
	Image         string `json:"image"`
	Nom_du_groupe string `json:"name"`
}

var home []Accueil

type Artists struct {
	Id_groupe     int      `json:"id"`
	Image         string   `json:"image"`
	Nom_du_groupe string   `json:"name"`
	Membres       []string `json:"members"`
	Creation      string   `json:"creationDate"`
	PremierAlbum  string   `json:"firstAlbum"`
	Locations     []string `json:"locations"`
	Dates         []string `json:"dates"`
}

var groupe []Artists

func HomePage(w http.ResponseWriter, r *http.Request) {
	// ici on récupère l'API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	// gestion d'erreur de la récup de l'API.
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// lecture de l'API
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &home)
	t, err := template.ParseFiles("./Static/HTML/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, home)

}

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./HTML/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, groupe)
}

func Artistes(w http.ResponseWriter, r *http.Request) {
	// On définit un chemin précis pour différencier tout le monde
	if r.URL.Path != "/Artiste" {
		http.Error(w, "404 not found", http.StatusNotFound)
		fmt.Println("404 link not found")
		return
	}

	// On récupère l'id pour la fonction de search
	id := r.URL.Query().Get("id")

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)

	if err != nil {
		http.Error(w, "500 Internal error server", http.StatusInternalServerError)
		fmt.Println("Erreur 500")
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		http.Error(w, "500 No data sent", http.StatusInternalServerError)
		fmt.Println("Erreur 500")
		return
	}

	// autrechoses = 1 Artiste
	var autrechoses Artists

	json.Unmarshal(responseData, &autrechoses)
	w.Header().Add("Content-Type", "application/json")
	w.Write(responseData)

	custTemplate, err := template.ParseFiles("./Static/HTML/artiste.html")

	if err != nil {
		http.Error(w, "500 no Data sent", http.StatusInternalServerError)
		fmt.Println("Erreur 500")
		return
	}

	err = custTemplate.Execute(w, var autrechose)

}
