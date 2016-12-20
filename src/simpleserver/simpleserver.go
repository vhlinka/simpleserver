//
// Copy/hack of sample code from "The Go Programming Language; Donovan & Kerighan"
//
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

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
