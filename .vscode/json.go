package main

// TimelineJSON represents user timeline
type TimelineJSON struct {
	Locations []Location
}

// ActivityWithConfidenceJSON from Google API
type ActivityWithConfidenceJSON struct {
	Type       string
	Confidence int
}

// ActivityJSON from Google API
type ActivityJSON struct {
	TimestampMs                string
	ActivityWithConfidenceJSON []ActivityWithConfidenceJSON
}

// Location from Google API
type Location struct {
	TimestampMs string
	LatitudeE7  int64
	LongitudeE7 int64
	Accuracy    int
	Activity    []ActivityJSON
}

type MapMarkerJSON struct {
	Name string
	Lat  float64
	Lng  float64
}

type InfectedPlaceJSON struct {
	MapMarker MapMarkerJSON
	Intensity int
}

type InfectedPlacesJSON struct {
	Places []InfectedPlaceJSON
}

type DurationJSON struct {
	StartTimestampMs string
	EndTimestampMs   string
}

type LatLangJSON struct {
	latitudeE7  int64
	LongitudeE7 int64
}

type ActivitiesJSON struct {
	ActivityType string
	probability  float64
}

type WaypointPathJSON struct {
	Waypoints []struct {
		LatE7 int64
		LngE7 int64
	}
}

type ActivitySegmentJSON struct {
	StartLocation LatLangJSON
	EndLocation   LatLangJSON
	Duration      DurationJSON
	Distance      int
	ActivityType  string
	Confidence    string
	Activities    []ActivitiesJSON
	WaypointPath  WaypointPathJSON
}

type LocJSON struct {
	LatitudeE7  int64
	LongitudeE7 int64
	PlaceId     string
	Address     string
	Name        string
	SourceInfo  struct {
		DeviceTag int64
	}
}

type PlaceVisitJSON struct {
	Location                LocJSON
	LocationConfidence      float64
	Duration                DurationJSON
	PlaceConfidence         string
	CenterLatE7             int64
	CenterLngE7             int64
	VisitConfidence         float64
	OtherCandidateLocations []struct {
		LatitudeE7         int64
		LongitudeE7        int64
		PlaceId            string
		LocationConfidence float64
	}
	EditConfirmationStatus string
	SimplifiedRawPath      struct {
		Points []struct {
			latE7          int64
			lngE7          int64
			TimestampMs    string
			AccuracyMeters int
		}
	}
}

type TimelineObjectJSON struct {
	ActivitySegment ActivitySegmentJSON
	PlaceVisit      PlaceVisitJSON
}

type RetroMovementJSON struct {
	TimelineObjects []TimelineObjectJSON
}
