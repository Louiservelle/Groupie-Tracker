package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func artist(w http.ResponseWriter, r *http.Request) {
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

	var artistes []Artist
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

func main() {

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	http.HandleFunc("/", homepage)
	http.HandleFunc("/artist", artist)

	log.Fatal(http.ListenAndServe(":80", nil))
}
