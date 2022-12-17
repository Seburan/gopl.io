// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// !+
func handleConn(c net.Conn) {

	var wg sync.WaitGroup // number of working go routines

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1) // add 1 worker
		go func(c net.Conn, shout string) {
			defer wg.Done() // -1 worker when go routine complete
			echo(c, shout, 1*time.Second)
		}(c, input.Text())
	}

	// closer wait for go routines complete before closing the connection
	go func() {
		wg.Wait()
		// NOTE: ignoring potential errors from input.Err()
		c.Close()
	}()

}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
