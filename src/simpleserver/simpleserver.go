//
// Copy/hack of sample code from "The Go Programming Language; Donovan & Kerighan"
//
//
//
// --- temp - for eval only:
// Google API key: AIzaSyCb6vGXygPRtFCePOoXpu221gl01ggBZqA
//
//
// Thank you to : http://www.kaihag.com/https-and-go/ for tips on the following https code changes
//
//
// for unsigned certs: openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
//
//
// Need new cert/key from cert from authority for non-dev env...
//
//
package main

import (
	"handlers"
	//	"log"
	"net/http"
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request.
	http.Redirect(w, r, "https://localhost:8001"+r.RequestURI, http.StatusMovedPermanently)
}

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

	// AwareAbility Gateway handlers

	http.HandleFunc("/SetToGatewayMode", handlers.EnterGatewayMode)
	http.HandleFunc("/ExitGatewayMode", handlers.ExitGatewayMode)

	//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//	log.Fatal(http.ListenAndServe(":8000", nil))
	////	go http.ListenAndServeTLS(":8001", "/media/vhlinka/AddedSpace/securityKeys/cert.pem", "/media/vhlinka/AddedSpace/securityKeys/key.pem", nil)
	go http.ListenAndServeTLS(":8001", "/media/vhlinka/AddedSpace/securityKeys/cert.pem", "/media/vhlinka/AddedSpace/securityKeys/key.pem", nil)
	// Start the HTTP server and redirect all incoming connections to HTTPS
	http.ListenAndServe(":8000", http.HandlerFunc(redirectToHttps))

	//	log.Fatal(http.ListenAndServeTLS(":8001", "/media/vhlinka/AddedSpace/securityKeys/cert.pem", "/media/vhlinka/AddedSpace/securityKeys/key.pem", nil))
}
