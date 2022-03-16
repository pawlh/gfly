package main

import (
	"strconv"
	"time"
)

type GoogleLocationHistory struct {
	Locations []Locations `json:"locations"`
}

type Locations struct {
	Timestamp        Timestamp     `json:"timestampMs"`
	Latitude         GPSCoordinate `json:"latitudeE7"`
	Longitude        GPSCoordinate `json:"longitudeE7"`
	Accuracy         int32         `json:"accuracy"`
	Altitude         int32         `Json:"altitude"`
	VerticalAccuracy int32         `json:"verticalAccuracy"`
	Activity         []Activity    `json:"activity"`
}

type Activity struct {
	TimestampMs Timestamp         `json:"timestampMs"`
	Activity    []ActivityDetails `json:"activity"`
}

type ActivityDetails struct {
	Type       string `json:"type"`
	Confidence int32  `json:"confidence"`
}

type Timestamp struct {
	Duration  time.Duration
	Timestamp time.Time
}

/* convert string timestamp to time */
func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1:][:len(str)-2]

	ms, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*ts = Timestamp(Timestamp{
		Duration:  time.Duration(ms) * time.Millisecond,
		Timestamp: time.Unix(0, ms*int64(time.Millisecond)),
	})
	return nil
}

type GPSCoordinate float32
