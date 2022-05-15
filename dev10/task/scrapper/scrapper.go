package scrapper

import (
	"dev10/common"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Scrapper struct{}

func New() Scrapper {
	return Scrapper{}
}

func (s *Scrapper) getLinks(url string) (common.Page, []string, error) {
	var (
		page common.Page
	)
	attachments := make([]string, 0)
	resp, err := http.Get(url)
	if err != nil {
		return common.Page{}, nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return common.Page{}, nil, err
	}

	page.URL = url

	f := func(n *html.Node) {
		for _, a := range n.Attr {
			if a.Key == "style" {
				if strings.Contains(a.Val, "url(") {

				}
			}
		}
	}
}

