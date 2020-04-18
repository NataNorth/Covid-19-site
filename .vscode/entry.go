package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func getInfectedPlacesHandler(w http.ResponseWriter, r *http.Request) {
	parentDir := getParentDir()
	timeline := readInfectedPeople(parentDir + "\\data\\infected.json")
	infectedPlaces := getInfectedPlaces(timeline)
	file, _ := json.MarshalIndent(infectedPlaces, "", " ")

	_ = ioutil.WriteFile(parentDir+"\\Data\\test.json", file, 0644)
	http.ServeFile(w, r, parentDir+"\\Data\\infected_places.json")
}

func uploadTimelineHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := req.FormFile("file")
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
	contents := buf.String()
	fmt.Println(contents)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	buf.Reset()
	// do something else
	// etc write header
	return
}

func getParentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(dir)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getInfectedPlacesHandler).Methods("GET")
	router.HandleFunc("/upload", uploadTimelineHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":12345", router))
}
