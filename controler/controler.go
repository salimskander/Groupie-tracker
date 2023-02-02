package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Accueil struct {
	Id_groupe     int    `json:"id"`
	Image         string `json:"image"`
	Nom_du_groupe string `json:"name"`
}

var home []Accueil

type Artist struct {
	Id_groupe     int      `json:"id"`
	Image         string   `json:"image"`
	Nom_du_groupe string   `json:"name"`
	Membres       []string `json:"members"`
	Creation      int      `json:"creationDate"`
	PremierAlbum  string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	Dates         string   `json:"dates"`
}

var artist Artist

type Relations struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	Id_groupe      int                 `json:"id"`
	Dates_location map[string][]string `json:"dates_location"`
}

var Concerts Relations

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

func Concert(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)

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

	json.Unmarshal(responseData, &Concerts)
	fmt.Println(Concerts)
	custTemplate, err := template.ParseFiles("./Static/HTML/concert.html")

	if err != nil {
		http.Error(w, "500 no Data sent", http.StatusInternalServerError)
		fmt.Println("Erreur 500")
		return
	}

	custTemplate.Execute(w, &Concerts)

}

func Recherche(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var Search []Artist
		var Result []Artist

		err = json.Unmarshal(responseData, &Search)
		if err != nil {
			panic(err)
		}
		input := r.FormValue("query")
		for _, artist := range Search {
			if input == strings.ToLower(artist.Nom_du_groupe) || input == artist.Nom_du_groupe {
				Result = append(Result, artist)
				continue
			}
			if input == strings.ToLower(artist.PremierAlbum) || input == artist.PremierAlbum {
				Result = append(Result, artist)
				continue
			}
			if input == strconv.Itoa(artist.Id_groupe) {
				Result = append(Result, artist)
				continue
			}
			if input == strconv.Itoa(artist.Creation) {
				Result = append(Result, artist)
				continue
			}
			for _, membres := range artist.Membres {
				if input == strings.ToLower(membres) || input == membres {
					Result = append(Result, artist)
					continue
				}
			}

		}
		t, err := template.ParseFiles("./Static/HTML/search.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, Result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
