// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix string = "http://";

func main() {
	for _, url := range os.Args[1:] {
		// Ex 1.7 : sanytize url
		if !strings.HasPrefix(url, httpPrefix) {
			url = "http://" + url;
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Ex 1.6 : copy directly to stdout
		// b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body);
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Status Code : %s\n", resp.Status);
		fmt.Printf("%d bytes read from %s\n", b, url);
	}
}

//!-
