package webscan

import (
	"fmt"
	"log"
)

//
// Simple function to extract all html links found in the page pointed to by the url
//
func crawl(url string) []string {
	fmt.Println(url)

	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
