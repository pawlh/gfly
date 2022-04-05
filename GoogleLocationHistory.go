package main

import (
    "strconv"
    "time"
)

type GoogleLocationHistory struct {
    Locations []Locations `json:"locations"`
}

type Locations struct {
    Timestamp        Timestamp     `json:"timestamp"`
    Latitude         GPSCoordinate `json:"latitudeE7"`
    Longitude        GPSCoordinate `json:"longitudeE7"`
    Accuracy         int32         `json:"accuracy"`
    Altitude         int32         `Json:"altitude"`
    VerticalAccuracy int32         `json:"verticalAccuracy"`
    Activity         []Activity    `json:"activity"`
}

type Activity struct {
    TimestampMs Timestamp         `json:"timestamp"`
    Activity    []ActivityDetails `json:"activity"`
}

type ActivityDetails struct {
    Type       string `json:"type"`
    Confidence int32  `json:"confidence"`
}

type Timestamp struct {
    time time.Time
    ms   int64
}

// UnmarshalJSON /* convert string timestamp to time */
func (ts *Timestamp) UnmarshalJSON(data []byte) error {
    str := string(data)
    str = str[1:][:len(str)-2]
    date, err := time.Parse("2006-01-02T15:04:05.999Z", str)
    if err != nil {
        return err
    }

    *ts = Timestamp{
        date,
        date.UnixMilli(),
    }
    return nil
}

type GPSCoordinate float32

func (c *GPSCoordinate) UnmarshalJSON(data []byte) error {
    str := string(data)
    // vi, err := strconv.ParseInt(str, 10, 64)
    // v := float64(vi / 10e4)
    v, err := strconv.ParseFloat(str, 64)
    if err != nil {
        return err
    }
    v /= 10e6

    *c = GPSCoordinate(v)

    return nil
}
