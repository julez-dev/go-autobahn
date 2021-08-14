package goautobahn

import (
	"strings"
	"time"
)

// CustomTime is a wrapper struct to implement a custom UnmarshalJSON method
type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	time, err := time.Parse("2006-01-02T15:04:05-0700", str)

	if err != nil {
		return err
	}

	t.Time = time

	return nil
}
