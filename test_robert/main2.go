package main

import (
    "encoding/json"
    "fmt"
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

func homepage(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("test.html")
	title := "groupie-Tracker"
	template.Execute(w, title)

}

func main() {

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

    for _, artist := range artistes {
        fmt.Println(artist.Name)

    } 

    }