package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const url = "0.beevik-ntp.pool.ntp.org"

func getTime() time.Time {
	tm, err := ntp.Time(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return tm
}
