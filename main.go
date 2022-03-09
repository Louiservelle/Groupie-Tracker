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

type Artisttest struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string //lien API
	ConcertDates string //lien API
	Relations    string //lien API
}
type relation struct {
	ID             int      `json:"id"`
	DatesLocations []string `json:"datesLocations"`
}
type Data struct {
	Artists []Artist
}

func getartist(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("artist.html")
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &artistes)

	Api := Data{
		Artists: artistes,
	}
	template.Execute(w, Api)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	template2, _ := template.ParseFiles("index.html")
	title := "groupie-Tracker"
	template2.Execute(w, title)
}

var artistes []Artist

func artistId(w http.ResponseWriter, r *http.Request) {

	template3, _ := template.ParseFiles("id.html")
	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)

	artistData := Artisttest{
		ID:           artistes[pathIDint-1].ID,
		Name:         artistes[pathIDint-1].Name,
		Image:        artistes[pathIDint-1].Image,
		Members:      artistes[pathIDint-1].Members,
		CreationDate: artistes[pathIDint-1].CreationDate,
		FirstAlbum:   artistes[pathIDint-1].FirstAlbum,
		Locations:    artistes[pathIDint-1].Locations,
		ConcertDates: artistes[pathIDint-1].ConcertDates,
		Relations:    artistes[pathIDint-1].Relations,
	}
	template3.Execute(w, artistData)
}
func getrelations(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https: //groupietrackers.herokuapp.com/api/relation/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &relation{})
}
func main() {

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	http.HandleFunc("/", homepage)
	http.HandleFunc("/artist", getartist)
	http.HandleFunc("/artist/", artistId)
	log.Fatal(http.ListenAndServe(":80", nil))

}
