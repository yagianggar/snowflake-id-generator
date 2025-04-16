package timeutil

import (
	"time"
)

func TimeToISO8601(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.RFC3339Nano)
}

func ParseISO8601(value string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// EndOfDay returns the very end of the given day.
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfDay returns the very end of the given day.
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
