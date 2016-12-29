package webscan

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document
//
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("ERROR: getting %s: %s", url, resp.Status)
	}

	fmt.Printf("response Proto : %s : ", resp.Proto)
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	// --- anonymous function - visitNode() ....
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // skip bad URLs
				}
				links = append(links, link.String())
			}
		}
	}

	// now call the anoymous funciton
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
