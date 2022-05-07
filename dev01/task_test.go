package main

import "testing"

func TestNtpTime(t *testing.T) {
	tests := []struct {
		name    string
		isValid bool
		host    string
	}{
		{
			"valid",
			true,
			"0.beevik-ntp.pool.ntp.org",
		},
		{
			"not valid",
			false,
			"host",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid == true {
				ntp, err := NtpTime(test.host)
				if err != nil {
					t.Fatal(err)
				} else if ntp == "" {
					t.Fatal()
				}
			} else {
				_, err := NtpTime(test.host)
				if err == nil {
					t.Fatal()
				}
			}
		})
	}
}
