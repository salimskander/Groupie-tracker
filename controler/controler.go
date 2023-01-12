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

type Artists struct {
	Id_groupe     int      `json:"id"`
	Image         string   `json:"image"`
	Nom_du_groupe string   `json:"name"`
	Membres       []string `json:"members"`
	Creation      string   `json:"creationDate"`
	PremierAlbum  string   `json:"firstAlbum"`
}

var groupe []Artists

func RecupApi(w http.ResponseWriter, r *http.Request) {

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

	json.Unmarshal(responseData, &groupe)
	fmt.Println(groupe)
	t, err := template.ParseFiles("./Static/HTML/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, groupe)

}

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./HTML/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, groupe)
}