package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var visitedPlaces VisitedPlaces

func getInfectedPlacesHandler() {
	parentDir := getParentDir()
	timeline := readInfectedPeople(parentDir + "\\Data\\infected.json")
	visitedPlaces = getVisitedPlaces(timeline)
	infectedPlaces := getInfectedPlaces(timeline, &visitedPlaces)
	file, _ := json.MarshalIndent(infectedPlaces, "", " ")

	_ = ioutil.WriteFile(parentDir+"\\website\\infected_places.json", file, 0644)
	//http.ServeFile(w, r, parentDir+"\\Data\\infected_places.json")
}

func getPlacesHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, getParentDir()+"\\website\\infected_places.json")
}

func uploadTimelineHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := req.FormFile("files[]")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	// Copy the file data to my buffer
	io.Copy(&buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	var retroMovement RetroMovementJSON
	err = json.Unmarshal(buf.Bytes(), &retroMovement)
	if err != nil {
		panic(err)
	}
	//hits := getHitsForPerson(retroMovement, &visitedPlaces)
	hits := 23

	text := []byte("You've been exposed " + strconv.Itoa(hits) + " times")
	err = ioutil.WriteFile(getParentDir()+"\\.vscode\\website\\Exposed.txt", text, 0644)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	buf.Reset()
	// do something else
	// etc write header
	return
}

func main() {

	// router := mux.NewRouter()
	// router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	// router.HandleFunc("/get", getInfectedPlacesHandler).Methods("GET")
	// router.HandleFunc("/upload", uploadTimelineHandler).Methods("POST")
	// log.Fatal(http.ListenAndServe(":8080", router))
	//getInfectedPlacesHandler()
	http.HandleFunc("/upload", uploadTimelineHandler)
	http.HandleFunc("/getplaces", getPlacesHandler)
	http.Handle("/", http.FileServer(http.Dir("website")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
