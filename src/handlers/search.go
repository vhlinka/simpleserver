package handlers

import (
	"net/http"
	"webscan"
)

func SearchGoogle(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	webscan.SearchGoogle(w, r)
}

func FindLinks(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	webscan.Findlinks(w)
}

func Crawler(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	webscan.CrawlLinks(w)
}

func Inspect(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	webscan.Inspect(w)
}

func Outline(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()
	webscan.Outline(w)
}

func FetchUrl(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	webscan.Fetch(w)
}
