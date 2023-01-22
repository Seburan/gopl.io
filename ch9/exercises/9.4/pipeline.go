// Exercise 9.4 (P280)
// Construct a pipeline that connects an arbitrary number of goroutines with
// channels. What is the maximum number of pipeline stages you can create with
// out running out of memory? How long does a value take to transit the entire pipeline

package pipeline

import (
	"errors"
)

func pipeline(stages int) (first chan int, last chan int, err error) {
	if stages < 1 {
		return nil, nil, errors.New("stages must me positive integer")
	}

	var in, out chan int
	in = make(chan int)
	first = in

	for i := 0; i < stages; i++ {
		out = make(chan int)
		go func(in, out chan int) {
			val := <-in
			out <- val
		}(in, out)
		in = out
	}

	last = out

	return first, last, nil
}
