package sitemap

import (
	"encoding/json"
	"io/ioutil"
)

type URLSitemap struct {
	Urls []struct {
		Url string `json:"url"`
	} `json:"urls"`
}

func (u *URLSitemap) Append(url string) {
	u.Urls = append(u.Urls, struct {
		Url string `json:"url"`
	}{Url: url})
}

func CreateSitemap(links []string) error {
	var total URLSitemap
	for _, v := range links {
		total.Append(v)
	}
	res, err := json.Marshal(total)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("wget/sitemap.json", res, 0644)
}
