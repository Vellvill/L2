package common

type Page struct {
	URL        string
	Canonincal string
	Links      []Links
}

type Links struct {
	Href string
}

var Root string

//PATH путь до папки
const PATH = "wget/"
