package handlers

//
// - implements a simple echo http handler - as per "the Go Programming Language"
//
import (
	"fmt"
	"graphics"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func EchoCounter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count : %d\n", count)
	mu.Unlock()

}

func EchoLissajous(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	graphics.Lissajuous(w)
}
