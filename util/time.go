package util

import (
	"time"
)

func HumanReadableDate(dateString string) string {
	// Parse the date string
	t, err := time.Parse(time.RFC3339Nano, dateString)
	if err != nil {
		return ""
	}

	// Format the time to a human-readable format
	humanReadable := t.Format("Monday, 02 January 2006, 03:04:05 PM")

	return humanReadable
}
