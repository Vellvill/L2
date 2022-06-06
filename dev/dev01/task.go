package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	cur, err := NtpTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cur)
}

func NtpTime(host string) (string, error) {
	ntpt, err := ntp.Time(host)
	if err != nil {
		return "", err
	}
	cur := ntpt.Format(time.UnixDate)
	return cur, nil
}
