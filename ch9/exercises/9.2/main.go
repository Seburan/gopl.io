// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+

// Exercise 9.2 : Rewrite the PopCount example from Section 2.6.2 so that it
// initializes the lookup table using sync.Once the first time it is needed.
// (Realistically, the cost of synchronization whould be prohibitive for a
//  small and highly optimized function like PopCount)

package popcount

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc []byte

// sync.Once
var loadLookupTableOnce sync.Once

func loadLookupTable() {
	pc = make([]byte, 256, 256)

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc[%d] = %d\n", i, pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	// only initialize the first time it is needed.
	loadLookupTableOnce.Do(loadLookupTable)

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-
