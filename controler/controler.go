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
}

type Relation struct {
	Dateslocations map[string][]string `json:"datesLocations"`
}

type Concertlist struct {
	Groupes   Accueil
	Relations Relation
}

var liste []Concertlist

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

	for id := 1; id < 53; id++ {

		var dataRelations Relation
		var dataGroupes Accueil
		var data Concertlist

		response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id) + "")
		if err != nil {
			http.Error(w, "500 Internal error server", http.StatusInternalServerError)
			fmt.Println("Erreur Impossible a Get")
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "500 No data sent", http.StatusInternalServerError)
			fmt.Println("Erreur ioutil.ReadAll")
			return
		}
		err = json.Unmarshal(responseData, &dataRelations)
		if err != nil {
			http.Error(w, "Umarhal problem : ---->", http.StatusInternalServerError)
			fmt.Println("Erreur Unmarshall dataRelation", err)
			return
		}

		demande, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(id) + "")
		if err != nil {
			http.Error(w, "500 Internal error server", http.StatusInternalServerError)
			fmt.Println("Erreur Impossible a Get artists")
			return
		}

		demandeData, err := ioutil.ReadAll(demande.Body)
		if err != nil {
			http.Error(w, "500 No data sent", http.StatusInternalServerError)
			fmt.Println("Erreur ioutil.ReadAllArtists")
			return
		}

		err = json.Unmarshal(demandeData, &dataGroupes)
		if err != nil {
			http.Error(w, "Umarshal problem : ->", http.StatusInternalServerError)
			fmt.Println("Erreur Unmarshall dataConcertlist", err)
			return
		}

		data.Groupes = dataGroupes
		data.Relations = dataRelations
		liste = append(liste, data)
		// fmt.Println(data)

	}

	custTemplate, err := template.ParseFiles("./Static/HTML/concert.html")

	if err != nil {
		http.Error(w, "500 no Data sent", http.StatusInternalServerError)
		fmt.Println("=Erreur 500", err)
		return
	}

	custTemplate.Execute(w, &liste)

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
