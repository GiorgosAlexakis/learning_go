//Exercise 5.2: Write a function to populate a mapping from element names—p, div, span, and
//so on—to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var countPerTagMap = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
	}
	countPerTag(doc)
	fmt.Println(countPerTagMap)
}

func countPerTag(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		countPerTagMap[n.Data]++
	}
	countPerTag(n.FirstChild)
	countPerTag(n.NextSibling)
	return
}
