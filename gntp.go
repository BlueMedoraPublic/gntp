package main

import (
	"fmt"
	"strconv"

	"github.com/beevik/ntp"
)

// version 0.1.0

// returns the current hour, minutes, seconds
// example: <hour> <min> <second> | 15 30 12
func main() {
	t, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println(strconv.Itoa(t.Hour()), strconv.Itoa(t.Minute()), strconv.Itoa(t.Second()))
	}
}
