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
