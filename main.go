package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Artist struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	StartYear  int      `json:"start_year"`
	FirstAlbum string   `json:"first_album"`
	Members    []string `json:"members"`
}

type Location struct {
	Locations []string `json:"locations"`
}

type Date struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	ArtistID  int      `json:"artist_id"`
	Locations []string `json:"locations"`
	Dates     []string `json:"dates"`
}

var (
	artists   []Artist
	locations []Location
	dates     []Date
	relations []Relation
)

func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func loadData() {
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}

	err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}

	err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		log.Fatalf("Error fetching dates: %v", err)
	}

	err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}
}

func main() {
	loadData()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artists", artistsHandler)
	http.HandleFunc("/locations", locationsHandler)
	http.HandleFunc("/dates", datesHandler)
	http.HandleFunc("/relation", relationHandler)
	http.HandleFunc("/refresh", refreshHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/artists.html"))
	tmpl.Execute(w, artists)
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/locations.html"))
	tmpl.Execute(w, locations)
}

func datesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dates.html"))
	tmpl.Execute(w, dates)
}

func relationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/relations.html"))
	tmpl.Execute(w, relations)
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	loadData()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
