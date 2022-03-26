// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const httpPrefix string = "http://";

func main() {
	start := time.Now()
	ch := make(chan string)
	var urlCounter int = 0;

	// Exercise 1.11 : read urls from file
	files := os.Args[1:]

	// process each file in argument
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetchall: %v\n", err)
			continue
		}
		defer f.Close();

		// process file as streem one by one
		input := bufio.NewScanner(f)
		for input.Scan() {
			url := input.Text();
			if !strings.HasPrefix(url, httpPrefix) {
				url = "http://" + url;
			}
			urlCounter++;
			go fetch(url, ch) // start a goroutine
		}

	}

	// wait for all urls to be processed
	for i := 1; i <= urlCounter; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
