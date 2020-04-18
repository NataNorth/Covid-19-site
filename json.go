package covid

type TimelineJSON struct {
	Locations []Location
}

type ActivityWithConfidenceJSON struct {
	Type string
	Confidence  int
}

type ActivityJSON struct {
	TimestampMs int64
	ActivityWithConfidenceJSON []ActivityWithConfidenceJSON
}

type Location struct {
	Timestamp int64
	LatitudeE7 int
	LongitudeE7 int
	Accuracy int
	Activity []ActivityJSON
}
