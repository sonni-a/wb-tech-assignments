package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	currentTime, err := getNTPTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(currentTime.Format(time.RFC3339))
}

func getNTPTime() (time.Time, error) {
	const ntpServer = "0.beevik-ntp.pool.ntp.org"
	t, err := ntp.Time(ntpServer)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
