// See page 134.

// Exercise 5.7 : Develop startElement and endElement into a general HTML
// pretty-printer. Print comment nodes, text nodes, and the attributes of each
// element (<a href='...'>). Use short forms like <img/> instead of <img></img>
// when an element has no children. Write a test to ensure that the output can
// be parsed successfully. (See Chapter 11.)

package main

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettify(t *testing.T) {

	// for testing purpose we replace the standard output with bytes buffer
	stdout = new(bytes.Buffer)

	// prettify("https://go.dev")
	prettify("https://www.w3schools.com/html/html_comments.asp")

	// once we have called the prettify function we want to check we can parse
	// its output
	_, err := html.Parse(bytes.NewReader(stdout.(*bytes.Buffer).Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
