package webscan

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func pageOutline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("ERROR: getting %s: %s", url, resp.Status)
	}

	fmt.Printf("response Proto : %s : ", resp.Proto)
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// now call the anoymous funciton
	forEachNode(doc, startElement, endElement)
	return nil
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

//
// print out the parsed html info ...
//
//
func inspectpage(url string) ([]string, error) {
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
	inspectNode := func(n *html.Node) {
		typeString := ""
		switch n.Type {
		case html.DoctypeNode:
			typeString = "DoctypeNode"
		case html.ErrorNode:
			typeString = "ErrorNode"
		case html.TextNode:
			typeString = "TextNode"
		case html.DocumentNode:
			typeString = "DocumentNode"
		case html.CommentNode:
			typeString = "CommentNode"
		case html.ElementNode:
			typeString = "ElementNode"
		default:
			typeString = "UNKNOWN Type"

		}
		links = append(links, "Node.Type: "+typeString)
		links = append(links, "Node.Data: "+n.Data)
		//		if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			links = append(links, "  a.Key: "+a.Key)
			links = append(links, "  a.Val: "+a.Val)
		}
		//		}
	}

	// now call the anoymous funciton
	forEachNode(doc, inspectNode, nil)
	return links, nil
}
