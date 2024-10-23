package models

import (
	"errors"
	"fmt"
	"time"
)

type JSONTime time.Time

const (
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(YYYYMMDDHHMMSS))
	return []byte(stamp), nil
}

func (t JSONTime) ConvertToYMD() string {
	return time.Time(t).Format(YYYYMMDDHHMMSS)
}

// UnmarshalJSON deserializes JSONTime from JSON
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	// Parse the incoming string to time
	s := string(b)
	if s == "null" {
		return nil
	}
	parsedTime, err := time.Parse(`"`+YYYYMMDDHHMMSS+`"`, s)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}
	// Set the parsed time to the JSONTime value
	*t = JSONTime(parsedTime)
	return nil
}
