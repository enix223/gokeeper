package datetime

import (
	"strconv"
	"time"
)

const (
	ns int64 = 1000000000
)

// FormatISOTime return a ISO format for given date
func FormatISOTime(date *time.Time) string {
	const form = "2006-01-02 15:04:05.00"
	return date.Format(form)
}

// UnixNanoTime convert unix nano timestamp to Time
func UnixNanoTime(ts int64) *time.Time {
	sec := ts / ns
	nano := ts - sec*ns
	t := time.Unix(sec, nano)
	return &t
}

// UnixTime convert unix timestamp to Time
func UnixTime(ts int64) *time.Time {
	t := time.Unix(ts, 0)
	return &t
}

// UnixNanoTimeString convert unix nano timestamp to Time
func UnixNanoTimeString(ts string) (*time.Time, error) {
	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil, err
	}
	return UnixNanoTime(t), nil
}

// UnixTimeString convert unix timestamp to Time
func UnixTimeString(ts string) (*time.Time, error) {
	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil, err
	}
	return UnixTime(t), nil
}
