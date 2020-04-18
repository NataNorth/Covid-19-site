package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	timeline := readInfectedPeople(filepath.Dir(dir) + "\\data\\infected.json")
	infectedPlaces := getInfectedPlaces(timeline)
	file, _ := json.MarshalIndent(infectedPlaces, "", " ")

	_ = ioutil.WriteFile(filepath.Dir(dir)+"\\Data\\test.json", file, 0644)

	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":12345", router))
}
