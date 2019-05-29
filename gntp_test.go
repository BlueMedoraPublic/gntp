package main

import (
    "strings"
    "strconv"
    "testing"
)

func TestGetTime(t *testing.T) {
    currentTime, err := GetTime()
	if err != nil {
		t.Errorf("got error while calling ntp.Time(): " + err.Error())
	}

    if len(currentTime) < 1 {
        t.Errorf("expected time string to be greater than length 0, got length: " + strconv.Itoa(len(currentTime)))
        return // return early, other tests will fail because of this
    }

    // split time string into an array delimited by
    // whitespace
    timeArray := strings.Fields(currentTime)
    if len(timeArray) != 3 {
        t.Errorf("expected time string to split into 3 when delimited by whitespace, got: " + strconv.Itoa(len(timeArray)))
        return // return early, other tests will fail because of this
    }

    for _, n := range timeArray {
        if len(n) != 2  {
            if len(n) != 1 {
                t.Errorf("testing length of hour, min, second. expected length to be 2, got: " + strconv.Itoa(len(n)))
            }
        }
    }
}
