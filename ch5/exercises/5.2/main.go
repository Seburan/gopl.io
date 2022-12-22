// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// exercise 5.2 : Write a function to populate a mapping from element names--
// p, div, span, and so on--to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countElements: %v\n", err)
		os.Exit(1)
	}

	elements := make(map[string]int)
	for key, value := range visit(elements, doc) {
		fmt.Printf("<%s> : %d\n", key, value)
	}
}

//!-main

// !+visit
// visit counts the number of each element in the HTML document tree
func visit(elements map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		elements[n.Data]++
		// for _, a := range n.Attr {
		// 	if a.Key == "href" {
		// 		links = append(links, a.Val)
		// 	}
		// }
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = visit(elements, c)
	}

	return elements
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
