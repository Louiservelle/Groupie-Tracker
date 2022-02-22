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

type Artist []struct {
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
	Name   string
	Img    string
	Global string
	Id     int
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

	var artistes Artist
	json.Unmarshal(responseData, &artistes)

	for i := 0; i < len(artistes); i++ {
		Api := Data{
			Name: artistes[i].Name,
			Img:  artistes[i].Image,
			Id:   artistes[i].ID,
		}
		template.Execute(w, Api)
	}
}

func main() {

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	http.HandleFunc("/", artist)

	log.Fatal(http.ListenAndServe(":80", nil))
}
