// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 134.

// Exercise 5.8 : Modify forEachNode so that the pre and post functions return
// a boolean result indicating whether to continue the traversal. Use it to
// write a function ElementByID with the following signature that finds the
// first HTML element with the specified id attribute. The function should
// stop the traversal as soon as a match is found.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stdin io.Reader = os.Stdin

func main() {

	searchID := os.Args[1] // get element to search from command line

	// OPTION 1 : get DOM from cUrl etc.
	// doc, err := html.Parse(os.Stdin)

	// OPTION 2 : get DOM directly from get
	resp, err := http.Get("https://go.dev")
	if err != nil {
		log.Fatal("cannot get url")
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	// END OF OPTION

	if err != nil {
		fmt.Fprintf(os.Stderr, "findElemenByID : %v\n", err)
		os.Exit(1)
	}
	element := findElementByID(doc, searchID)
	if element != nil {
		fmt.Fprintf(stdout, "Found element with id=%s : %v\n", searchID, element)
	} else {
		fmt.Fprintf(stdout, "NOT Found element with id=%s\n", searchID)
	}

}

func findElementByID(doc *html.Node, id string) *html.Node {

	//!+call
	element := forEachNode(doc, id, startElement, nil)
	//!-call

	return element
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {

	if pre != nil {
		ok := pre(n, id)
		if ok == true {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		element := forEachNode(c, id, pre, post)
		if element != nil {
			//fmt.Fprintf(stdout, "FOUND id=%s : %v\n", id, element)
			return element
		}
	}

	if post != nil {
		post(n, id)
	}

	return nil
}

//!-forEachNode

// !+startend
var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" {
				// fmt.Fprintf(stdout, "*** id=%s : %v\n", attr.Val, n)
			}
			if attr.Key == "id" && attr.Val == id {
				//fmt.Fprintf(stdout, "FOUND id=%s : %v\n", attr.Val, n)
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	// if n.Type == html.ElementNode {
	// 	depth--
	// 	fmt.Fprintf(stdout, "%*s</%s>\n", depth*2, "", n.Data)
	// }

	return false
}

//!-startend
