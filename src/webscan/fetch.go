package webscan

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(w io.Writer) {
	url := "https://www.rev1ventures.com/about/"

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(w, "ERROR: getting %s: %s", url, resp.Status)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "ERROR: reading %s: %v", url, err)
		return
	}

	fmt.Fprintf(w, "Contains = %v\n", bytes.Contains(bytes.ToLower(b), []byte("vasil")))

}
