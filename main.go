package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Data struct {
	Artists []Artist
}

type Datelocations struct {
	ID        int                 `json:"id"`
	Locations map[string][]string `json:"datesLocations"`
}

//----------------------------------------------------------------------------------------------------------------\\
//-----------------------------------------------FUNCTION-ARTIST--------------------------------------------------\\
//----------------------------------------------------------------------------------------------------------------\\

func getartists(w http.ResponseWriter, r *http.Request) {

	template, _ := template.ParseFiles("artist.html")

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &Artistes)

	Api := Data{
		Artists: Artistes,
	}
	template.Execute(w, Api)
}

var Artistes []Artist

func artistId(w http.ResponseWriter, r *http.Request) {
	template2, _ := template.ParseFiles("id.html")
	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	var gomerde Datelocations

	artistData := Artist{
		ID:           Artistes[pathIDint-1].ID,
		Name:         Artistes[pathIDint-1].Name,
		Image:        Artistes[pathIDint-1].Image,
		Members:      Artistes[pathIDint-1].Members,
		CreationDate: Artistes[pathIDint-1].CreationDate,
		FirstAlbum:   Artistes[pathIDint-1].FirstAlbum,
		Locations:    Artistes[pathIDint-1].Locations,
		ConcertDates: Artistes[pathIDint-1].ConcertDates,
		Relations:    Artistes[pathIDint-1].Relations,
	}
	resp, err := http.Get(artistData.Relations)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}
	json.Unmarshal(body, &gomerde)

	M := map[string]interface{}{
		"Artiste":  artistData,
		"Relation": gomerde,
	}
	template2.Execute(w, M)
}

//----------------------------------------------------------------------------------------------------------------\\
//-----------------------------------------------MAIN-------------------------------------------------------------\\
//----------------------------------------------------------------------------------------------------------------\\
func main() {
	//Declaration css path
	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	//Link path
	http.HandleFunc("/artist", getartists)
	http.HandleFunc("/artist/", artistId)
	http.HandleFunc("/", homepage)
	//Listening port 80
	log.Fatal(http.ListenAndServe(":80", nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	template3, _ := template.ParseFiles("index.html")
	template3.Execute(w, r)
}
