package models

import (
	"errors"
	"fmt"
	"time"
)

type JSONTime time.Time

const (
	YYYYMMDD = "2006-01-02"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	// do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(YYYYMMDD))
	return []byte(stamp), nil
}

func (t JSONTime) ConvertToYMD() string {
	return time.Time(t).Format(YYYYMMDD)
}

// UnmarshalJSON deserializes JSONTime from JSON
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	// Parse the incoming string to time
	s := string(b)
	if s == "null" {
		return nil
	}
	parsedTime, err := time.Parse(`"`+YYYYMMDD+`"`, s)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}
	// Set the parsed time to the JSONTime value
	*t = JSONTime(parsedTime)
	return nil
}
