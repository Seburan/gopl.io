// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountByPreAllocation8bits returns the population count (number of set bits) of x.
func PopCountByPreAllocation8bits(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}


// -- Alternative implementations --

// Exercise 2.3 Rewrite PopCount to use a loop and compare performances
func PopCountByPreAllocation8bitsLoop(x uint64) int {

	var count int = 0;

	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))]);
	}

	return count;
}

// Exercise 2.4 Count bits by shifting its argument through 64 bit positions testing the rightmost bit each time.
func PopCountByShiftingRight(x uint64) int {
	var count int = 0;

	// loop over each bit (shift) and check if 1, if so add 1 to counter
	for i := 0; i < 64; i++ {
		count += int(byte((x>>i)&1));
	}

	return count;
}

func PopCountByShiftingLeft(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}


func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}


// Exercise 2.5 : count bits by using x & (x-1) clearing the rightmost non-zero bit of x
func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}


//!-
