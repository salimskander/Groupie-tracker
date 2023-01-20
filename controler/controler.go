package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	Creation      int      `json:"creationDate"`
	PremierAlbum  string   `json:"firstAlbum"`
	Locations     []string `json:"locations"`
	Dates         []string `json:"dates"`
}

var artist Artists

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
	t.Execute(w, home)
}

func Artiste(w http.ResponseWriter, r *http.Request) {

	// On récupère l'id
	id := r.URL.Query().Get("id")

	//if !id -> redirect

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

	json.Unmarshal(responseData, &artist)
	fmt.Println(responseData)
	fmt.Println(&artist)
	custTemplate, err := template.ParseFiles("./Static/HTML/artiste.html")

	if err != nil {
		http.Error(w, "500 no Data sent", http.StatusInternalServerError)
		fmt.Println("Erreur 500")
		return
	}

	custTemplate.Execute(w, &artist)

}

func Recherche(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Recherche" {
		http.Error(w, "404 not found", http.StatusNotFound)
		fmt.Println("404 link not found")
		return
	}

	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	id := params.Get("id")
	page := params.Get("/artists/" + id)
	if page == "" {
		page = "1"
	}

	fmt.Println("Artiste : ", page)

}
