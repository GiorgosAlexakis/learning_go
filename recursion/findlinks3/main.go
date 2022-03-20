// Exercise 5.4: Extend the visit function(findlinks1) so that it extracts other kinds of links from
// the document, such as images, scripts, and style sheets
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var mapping = map[string]string{"a": "href", "img": "src", "script": "src", "link": "href"}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
	}
	mapLinks := make(map[string][]string)
	visit(mapLinks, doc)
	fmt.Println(mapLinks["link"])
}

// visit appends to links each link found in n and returns the result.
func visit(links map[string][]string, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		if _, ok := mapping[n.Data]; ok {
			for _, a := range n.Attr {
				if a.Key == mapping[n.Data] {
					links[n.Data] = append(links[n.Data], a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
	return
}
