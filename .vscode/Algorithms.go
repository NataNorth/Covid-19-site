package main

import (
	"context"
	"log"
	"time"

	"googlemaps.github.io/maps"
)

// DaInMs is a day in ms
const DayInMS = 86400000

// GCThreshold is a garbage collection threshold
const GCThreshold = 3 * DayInMS

// Place is a simple place representation
type Place struct {
	Name    string
	Address string
	Latlng  maps.LatLng
}

// VisitedPlaces has a method to get place infection intensity
type VisitedPlaces struct {
	table map[Place]map[int64]int
}

func (vp VisitedPlaces) getIntensity(p Place) int {
	count := 0
	for _, timestampVisit := range vp.table[p] {
		count += timestampVisit
	}
	return count
}

func expired(ms int64) bool {
	now := time.Now()
	umillisec := now.UnixNano() / 1000000
	if ms < umillisec-GCThreshold {
		return true
	}
	return false
}

func (vp *VisitedPlaces) collectGarbage() {
	for place := range vp.table {
		for timestamp := range vp.table[place] {
			if expired(timestamp) {
				delete(vp.table[place], timestamp)
			}
		}
	}
}

var visitedPlaces VisitedPlaces

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
	ll.Lat = float64(lat / 1e7)
	ll.Lng = float64(lang / 1e7)
	return ll
}

func treatPoint(timestamp int64, lat int64, lng int64, c *maps.Client) {

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
		if !place.PermanentlyClosed {
			formattedPlace := Place{place.Name, place.FormattedAddress, ll}
			visitedPlaces.table[formattedPlace][timestamp]++
			break
		}
	}
}

func getInfectedPlaces(timeline TimelineJSON) InfectedPlacesJSON {
	c := getClient()
	for _, place := range timeline.Locations {
		treatPoint(place.Timestamp, place.LatitudeE7, place.LongitudeE7, c)
	}
	var infectedPlaces InfectedPlacesJSON
	for place := range visitedPlaces.table {
		infectedPlace := InfectedPlaceJSON{MapMarkerJSON{place.Name, place.Address, place.Latlng.Lat, place.Latlng.Lng}, visitedPlaces.getIntensity(place)}
		infectedPlaces.Places = append(infectedPlaces.Places, infectedPlace)
	}
	return infectedPlaces
}
