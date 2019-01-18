package gokeeper

import (
	"testing"
	"time"
)

func TestDateISOFormat(t *testing.T) {
	// Mon, 02 Jan 2006 15:04:05 MST
	d, _ := time.Parse(time.RFC1123, time.RFC1123)
	val := "2006-01-02 15:04:05.00"
	res := FormatISOTime(&d)
	if res != val {
		t.Errorf(expectValue, val, res)
	}
}
