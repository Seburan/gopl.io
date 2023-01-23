// Exercise 9.5 (P281)
// Write a program with two gotoutines that send messages back and forth over
// two unbuffered channels in ping-pong fashion. How many communications per second can the program sustain?

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("program start")

	// ping pong channel
	var ping, pong chan int
	ping = make(chan int)
	pong = make(chan int)

	// end ping pong
	done := make(chan struct{})

	// counter # of communications
	var counter int

	// send from ping
	go func() {
		for val := range ping {
			select {
			case <-done:
				return
			case pong <- val:
				counter++
			}
		}
	}()

	// send from pong
	go func() {
		for val := range pong {
			select {
			case <-done:
				return
			case ping <- val:
			}
		}
	}()

	// start time
	start := time.Now()

	// send the first ball
	ping <- 1

	// sleep for 1000 sec
	time.Sleep(time.Millisecond * 1000)

	// stop go routines
	close(done)
	close(ping)
	close(pong)

	fmt.Printf("Pingpong achieved %d communications in %d ms\n ", counter, time.Since(start).Milliseconds())
}
