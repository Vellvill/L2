package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	url string
)

func init() {
	flag.StringVar(&url, "url", "", "use for setting url")
}

func main() {
	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println()
	}
	defer resp.Body.Close()

}
