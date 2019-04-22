package datetime

import (
	"math"
	"time"
)

// FormatISOTime return a ISO format for given date
func FormatISOTime(date *time.Time) string {
	const form = "2006-01-02 15:04:05.00"
	return date.Format(form)
}

// UnixNanoTime convert unix nano timestamp to Time
func UnixNanoTime(ts int64) *time.Time {
	ns := int64(math.Pow(10, 9))
	sec := ts / ns
	nano := ts - sec*ns
	t := time.Unix(sec, nano)
	return &t
}
