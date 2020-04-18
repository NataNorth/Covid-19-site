package covid

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func readJSON(filename string) (*LocationData, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var ld LocationData
	err = json.Unmarshal(b, &ld)
	if err != nil {
		return nil, err
	}
	return &ld, nil
}

func main() {
	ld, err := readJSON("data/data.json")
	if err != nil {
		log.Fatal(err)
	}
	ld.Data[0].Lat = 1
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	log.Fatal(http.ListenAndServe(":12345", router))
}
