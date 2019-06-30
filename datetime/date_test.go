package datetime

import (
	"strconv"
	"testing"
	"time"
)

func TestDateISOFormat(t *testing.T) {
	// Mon, 02 Jan 2006 15:04:05 MST
	d, _ := time.Parse(time.RFC1123, time.RFC1123)
	val := "2006-01-02 15:04:05.00"
	res := FormatISOTime(&d)
	if res != val {
		t.Errorf("expect: %v, got: %v", val, res)
	}
}

func TestUnixNanoTime(t *testing.T) {
	tn := time.Now()
	ts := tn.UnixNano()
	tc := UnixNanoTime(ts)
	if tn.Sub(*tc) != 0 {
		t.Errorf("exp: %v, got: %v", tn, tc)
	}
}

func TestUnixTime(t *testing.T) {
	tn := time.Now()
	ts := tn.Unix()
	tc := UnixTime(ts)
	if tn.Sub(*tc) > 1*time.Second {
		t.Errorf("exp: %v, got: %v", tn, tc)
	}
}

func TestUnixNanoTimeString(t *testing.T) {
	tn := time.Now()
	ts := tn.UnixNano()
	tc, err := UnixNanoTimeString(strconv.FormatInt(ts, 10))
	if err != nil {
		t.Error(err)
	}
	if tn.Sub(*tc) != 0 {
		t.Errorf("exp: %v, got: %v", tn, tc)
	}
}

func TestUnixTimeString(t *testing.T) {
	tn := time.Now()
	ts := tn.Unix()
	tc, err := UnixTimeString(strconv.FormatInt(ts, 10))
	if err != nil {
		t.Error(err)
	}
	if tn.Sub(*tc) > 1*time.Second {
		t.Errorf("exp: %v, got: %v", tn, tc)
	}
}
func TestUnixNanoTimeStringFailed(t *testing.T) {
	_, err := UnixNanoTimeString("abc")
	if err == nil {
		t.Fail()
	}
}

func TestUnixTimeStringFailed(t *testing.T) {
	_, err := UnixTimeString("abc")
	if err == nil {
		t.Fail()
	}
}
