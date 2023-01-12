package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RecupApi(w http.ResponseWriter, r *http.Request) {

	// ici on récupère l'API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")

	// gestion d'erreur de la récup de l'API.
	if err != nil {
		http.Error(w, "500 - Données non récupérées", http.StatusInternalServerError)
		fmt.Println("[SERVER_ALERT] - 500 : Internal server error")
		return
	}
	// lecture de l'API
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "500 - Aucunes données envoyées", http.StatusInternalServerError)
		fmt.Println("[SERVER_ALERT] - 500 : Internal server error")
	}
	fmt.Println(string(responseData))

	var StructuresObjet Artists
	json.Unmarshal(responseData, &StructuresObjet)
}

// création de structures pour les différentes données

type Artists struct {
	Id_groupe     int      `json:"id"`
	Image         string   `json:"image"`
	Nom_du_groupe string   `json:"name"`
	Membres       []string `json:"members"`
	Creation      string   `json:"creationDate"`
	PremierAlbum  string   `json:"firstAlbum"`
}

type Index struct {
	Index []int `json:"index"`
}

type Lieux struct {
	Id    int         `json:"id"`
	Lieux []Locations `json:"locations"`
	Dates string      `json:"dates"`
}

type Locations struct {
	Lieux string `json:"locations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id        int      `json:"id"`
	Date_Lieu string   `json:"datesLocations"`
	Lieu_date []string `json:""`
}
