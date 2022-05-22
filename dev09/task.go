package main

import (
	"L2/dev09/client"
	"flag"
	"log"
)

var (
	timeout int
	ip      string
	port    string
)

func init() {
	flag.IntVar(&timeout, "timeout", 0, "connection timeout")
	flag.StringVar(&ip, "ip", "", "ip to connect")
	flag.StringVar(&port, "port", "", "port to connect")
}

func main() {
	flag.Parse()
	tel := client.NewTelnet(port, ip, timeout)
	err := tel.Dial()
	if err != nil {
		log.Fatal(err)
	}
}
