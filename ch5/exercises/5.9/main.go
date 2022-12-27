// Exercise 5.9 : Write a function
// expand(s string, f func(string) string) string
// that replaces each substring "$foo" within s by the text returned by f("foo").

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var stdout io.Writer = os.Stdout
var stdin io.Reader = os.Stdin

func main() {

	// define a function value
	var myFunction func(string) string
	myFunction = func(s string) string {
		return s + "bar"
	}

	for _, arg := range os.Args[1:] {
		fmt.Fprintf(stdout, "%s ", expand(arg, myFunction))
	}
	fmt.Println()

	return
}

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$foo", f("foo"))
}
