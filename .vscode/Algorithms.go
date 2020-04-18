package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"googlemaps.github.io/maps"
)

// DaInMs is a day in ms
const DayInMS = 86400000

// GCThreshold is a garbage collection threshold
const GCThreshold = 3 * DayInMS

// Place is a simple place representation
type PlaceDetails struct {
	TimestampMs string
	Latlng      maps.LatLng
}

// VisitedPlaces has a method to get place infection intensity
type VisitedPlaces struct {
	placeVisitsTable  map[string]map[string]int
	placeDetailsTable map[string]maps.LatLng
}

func (vp VisitedPlaces) getIntensity(place string) int {
	count := 0
	for _, timestampVisit := range vp.placeVisitsTable[place] {
		count += timestampVisit
	}
	return count
}

func expired(ms string) bool {
	now := time.Now()
	umillisec := now.UnixNano() / 1000000
	n, _ := strconv.Atoi(ms)
	if int64(n) < umillisec-GCThreshold {
		return true
	}
	return false
}

func (vp *VisitedPlaces) collectGarbage() {
	for place, timestamps := range vp.placeVisitsTable {
		for timestamp := range timestamps {
			if expired(timestamp) {
				delete(vp.placeVisitsTable[place], timestamp)
				delete(vp.placeDetailsTable, place)
			}
		}
	}
}

func getClient() *maps.Client {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCo2JSgqGWWDTEEUl6gv1Ys2Kj2FyuS630"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	return c
}
func normalizeLatLang(lat int64, lang int64) maps.LatLng {
	if lat > 900000000 {
		lat -= 4294967296
	}
	if lang > 1800000000 {
		lang -= 4294967296
	}
	var ll maps.LatLng
	ll.Lat = float64(lat) / 1e7
	ll.Lng = float64(lang) / 1e7
	return ll
}

func filterTypes(types []string) bool {
	for _, t := range types {
		if t == "locality" || t == "political" || t == "country" || t == "continent" || t == "colloquial_area" || t == "archipelago" || t == "postal_code" || t == "street_address" || t == "natural_feature" || t == "geocode" || t == "intersection" || t == "town_square" || t == "neighborhood" || t == "street_address" {
			return true
		}
	}
	return false
}

func treatPoint(timestamp string, lat int64, lng int64, localVisitedPlaces *VisitedPlaces, c *maps.Client) {
	if expired(timestamp) {
		return
	}
	ll := normalizeLatLang(lat, lng)
	r := &maps.NearbySearchRequest{
		Location: &ll,
		Radius:   10,
	}

	response, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatal(err)
	}
	for _, place := range response.Results {
		if filterTypes(place.Types) {
			continue
		}
		if !place.PermanentlyClosed {
			log.Print("Adding place: " + place.Name)
			_, ok := localVisitedPlaces.placeVisitsTable[place.Name]
			if !ok {
				localVisitedPlaces.placeVisitsTable[place.Name] = make(map[string]int)
			}
			localVisitedPlaces.placeVisitsTable[place.Name][timestamp]++
			localVisitedPlaces.placeDetailsTable[place.Name] = ll
		}
	}
}

func getVisitedPlaces(timeline TimelineJSON) VisitedPlaces {
	c := getClient()
	visitedPlaces := VisitedPlaces{make(map[string]map[string]int), make(map[string]maps.LatLng)}
	for _, place := range timeline.Locations {
		treatPoint(place.TimestampMs, place.LatitudeE7, place.LongitudeE7, &visitedPlaces, c)
	}
	return visitedPlaces
}

func getInfectedPlaces(timeline TimelineJSON, visitedPlaces *VisitedPlaces) InfectedPlacesJSON {
	var infectedPlaces InfectedPlacesJSON
	for placeName := range visitedPlaces.placeVisitsTable {
		placell := visitedPlaces.placeDetailsTable[placeName]
		infectedPlace := InfectedPlaceJSON{MapMarkerJSON{placeName, placell.Lat, placell.Lng}, visitedPlaces.getIntensity(placeName)}
		infectedPlaces.Places = append(infectedPlaces.Places, infectedPlace)
	}
	return infectedPlaces
}

func readInfectedPeople(filename string) TimelineJSON {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var timeline TimelineJSON
	err2 := json.Unmarshal(b, &timeline)
	if err2 != nil {
		log.Fatal(err2)
	}
	return timeline
}

func getHitsForPerson(movement RetroMovementJSON, visitedPlaces *VisitedPlaces) int {
	hits := 0
	for _, timelineObject := range movement.TimelineObjects {
		if !expired(timelineObject.PlaceVisit.Duration.EndTimestampMs) {
			_, ok := visitedPlaces.placeDetailsTable[timelineObject.PlaceVisit.Location.Name]
			if ok {
				hits++
			}
		}
	}
	return hits
}
