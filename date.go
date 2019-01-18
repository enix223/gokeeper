package gokeeper

import "time"

// FormatISOTime return a ISO format for given date
func FormatISOTime(date *time.Time) string {
	const form = "2006-01-02 15:04:05.00"
	return date.Format(form)
}
