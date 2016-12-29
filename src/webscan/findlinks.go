package webscan

import (
	"fmt"
	"io"
	"log"
)

var out io.Writer

func Findlinks(w io.Writer) {

	out = w
	printHtmlHeader()

	// - scan a url for links
	var worklist []string

	url := "https://www.rev1ventures.com/about/"
	fmt.Fprintf(out, "<P>Scanning for links - URL: %s</P>\n", url)

	items, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	worklist = append(worklist, items...)

	for _, item := range worklist {
		fmt.Fprintf(out, "Link: %s<br>\n", item)
	}

	printHtmlFooter()

}

//
// function that will extract html links by crawling through from a starting point
//
func CrawlLinks(w io.Writer) {
	printHtmlHeader()

	var urlList []string
	urlList = append(urlList, "https://www.rev1ventures.com/about/")
	fmt.Fprintf(out, "<P>CRAWLING for links - Starting with URL: %s</P>\n", urlList[0])

	items := breadthFirst(crawl, urlList)
	for _, item := range items {
		fmt.Fprintf(out, "Link: %s<br>\n", item)
	}

	printHtmlFooter()
}

func printHtmlHeader() {
	fmt.Fprintf(out, "<!DOCTYPE html>\n<html>\n<head>\n<title>HTML Link Extractor</title>\n</head>\n<Body>\n")
}

func printHtmlFooter() {
	fmt.Fprintf(out, "</Body>\n</html>\n")
}

// used to track the depth of the html node traversal
var depth int

func Inspect(w io.Writer) {

	out = w
	printHtmlHeader()

	// - scan a url for links
	var worklist []string

	url := "https://www.rev1ventures.com/about/"
	fmt.Fprintf(out, "<P>Scanning for links - URL: %s</P>\n", url)

	items, err := inspectpage(url)
	if err != nil {
		log.Print(err)
	}
	worklist = append(worklist, items...)

	for _, item := range worklist {
		fmt.Fprintf(out, "%s<br>\n", item)
	}

	printHtmlFooter()

}

func Outline(w io.Writer) {

	out = w
	printHtmlHeader()

	url := "https://www.rev1ventures.com/about/"
	fmt.Fprintf(out, "<P>Scanning for links - URL: %s</P>\n", url)

	err := pageOutline(url)
	if err != nil {
		log.Print(err)
	}

	printHtmlFooter()

}
