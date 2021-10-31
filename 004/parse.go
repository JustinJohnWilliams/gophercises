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
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	// NOTE: implementing depth first search
	// for loop assigns c to first child, if it's not nil, do the thing
	// on the last element change c to next sibling?? not sure how for loops work in go. wtf
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}

}
