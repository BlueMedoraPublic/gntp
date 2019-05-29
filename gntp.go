package main

import (
	"os"
	"fmt"
	"time"
	"strconv"

	"github.com/beevik/ntp"
)

// version 0.1.1

const ntpServer string = "pool.ntp.org"

// print the time and exit. Print errors to Stderr and exit
// with code 1
func main() {
	t, err := GetTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(t)
}

// returns the current hour, minutes, seconds
func GetTime() (string, error) {
	t, err := ntp.Time(ntpServer)
	if err != nil {
		return "", err
	}

	// convert the time to a string and return
	return timeToString(t), nil
}

// convertTime takes a time.Time object and converts the
// hour, minutes, and seconds into a string
// example: <hour> <min> <second> | 15 30 12
func timeToString(t time.Time) string {
	return strconv.Itoa(t.Hour()) + " " + strconv.Itoa(t.Minute()) + " " + strconv.Itoa(t.Second())
}
