//
// Copy/hack of sample code from "The Go Programming Language; Donovan & Kerighan"
//
//
//
// --- temp - for eval only:
// Google API key: AIzaSyCb6vGXygPRtFCePOoXpu221gl01ggBZqA
//
package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.EchoHandler)
	http.HandleFunc("/count", handlers.EchoCounter)
	http.HandleFunc("/lissajous", handlers.EchoLissajous)
	http.HandleFunc("/mandelbrot", handlers.EchoMandelbrot)
	http.HandleFunc("/search", handlers.SearchGoogle)
	http.HandleFunc("/findlinks", handlers.FindLinks)
	http.HandleFunc("/inspect", handlers.Inspect)
	http.HandleFunc("/fetch", handlers.FetchUrl)
	http.HandleFunc("/crawl", handlers.Crawler)
	http.HandleFunc("/outline", handlers.Outline)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
