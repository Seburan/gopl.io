// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 134.

// Exercise 5.7 : Develop startElement and endElement into a general HTML
// pretty-printer. Print comment nodes, text nodes, and the attributes of each
// element (<a href='...'>). Use short forms like <img/> instead of <img></img>
// when an element has no children. Write a test to ensure that the output can
// be parsed successfully. (See Chapter 11.)

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stdin io.Reader = os.Stdin

func main() {
	for _, url := range os.Args[1:] {
		prettify(url)
	}
}

func prettify(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

//!-forEachNode

// !+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attributes strings.Builder
		for _, attr := range n.Attr {
			fmt.Fprintf(&attributes, " %s=\"%s\"", attr.Key, attr.Val)
		}
		fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attributes.String())
		depth++
	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			fmt.Fprintf(stdout, "%*s%s\n", depth*2, "", text)

		}
	} else if n.Type == html.CommentNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			fmt.Fprintf(stdout, "%*s<!-- %s -->\n", depth*2, "", text)
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(stdout, "%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
