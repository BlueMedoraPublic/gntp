package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"flag"

	"github.com/beevik/ntp"
)

const version string   = "0.1.1"
const ntpServer string = "pool.ntp.org"

// print the time and exit. Print errors to Stderr and exit
// with code 1
func main() {
	var printVersion bool
	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
	flag.Parse()

	if printVersion == true {
		fmt.Println(version)
		os.Exit(0)
	}

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
