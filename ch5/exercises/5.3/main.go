// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// exercise 5.3 : Write a function to print the contents of all text nodes.
// in an HTML document tree. Do not descend into <script> or <style> elements,
// since their contents are not visible in a web browser.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findContent: %v\n", err)
		os.Exit(1)
	}

	for i, content := range visit(nil, doc) {
		fmt.Println(i, content)
	}
}

//!-main

// !+visit
// visit finds the content of all text nodes in the HTML document tree
func visit(contents []string, n *html.Node) []string {

	if n.Type == html.TextNode {
		// ignore text under <script> and <style> nodes
		switch n.Parent.Data {
		case "script", "style":
			return contents
		}

		// most of the data are blank line or whitespaces
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			contents = append(contents, text)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = visit(contents, c)
	}

	return contents
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
