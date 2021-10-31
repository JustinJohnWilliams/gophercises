package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (a <href="" />) in HTML
type Link struct {
	Href string
	Text string
}

// Parse will take in HTML doc and return a slice of links:w
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}

	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// NOTE: the linkNodes returns a slice, the `...` returns individual elements of the slice
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}