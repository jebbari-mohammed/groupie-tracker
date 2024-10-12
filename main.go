package main

import (
	"fmt"
	"net/http"

	groupietracker "grptracker/functions"
)

func main() {
	fmt.Println("Server is running on http://localhost:8080/")
	http.HandleFunc("/", groupietracker.HomeHandler)
	http.HandleFunc("/artists", groupietracker.ArtistsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
