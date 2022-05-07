package main

import (
	"github.com/beevik/ntp"
	"time"
)

func main() {
	NtpTime("0.beevik-ntp.pool.ntp.org")
}

func NtpTime(host string) (string, error) {
	ntpt, err := ntp.Time(host)
	if err != nil {
		return "", err
	}
	cur := ntpt.Format(time.UnixDate)
	return cur, nil
}
