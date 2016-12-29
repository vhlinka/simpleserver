package webscan

import (
	"sort"
)

//
// function to visit each link contained in the list of urls passed in worklist
// return a sorted string with all of the urls visited
//
//
func breadthFirst(f func(item string) []string, worklist []string) []string {
	seen := make(map[string]bool) // track which links have been processed
	for len(worklist) > 0 {
		items := worklist // let items hold the current worklist batch
		worklist = nil    // ready the worklist for the next batch (from crawler)
		for _, item := range items {
			if !seen[item] { // only visit the url if it has not been processed at least once
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
	var urls []string
	for url := range seen {
		urls = append(urls, url)
	}
	sort.Strings(urls)
	return urls
}
