// Exercise 2.3 - 2.4 - 2.5
// Rewrite PopCount to use a loop and compare performances

package main

import (
	"fmt"

	"gopl.io/ch2/exercises/popcount"
)

// pc[i] is the population count of i.
var pc [256]byte

func main() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("PopCount of %d (%b) = %d\n", i, i, pc[i] )
	}

	var v uint64 = 0x1234567890ABCDEF
	fmt.Printf("decimal : %d hex : %x  binary : %b\n", v, v, v)
	fmt.Printf("popcount.PopCountByPreAllocation8bits(%x) = %d\n", v, popcount.PopCountByPreAllocation8bits(v))
	fmt.Printf("popcount.PopCountByPreAllocation8bitsLoop(%x) = %d\n", v, popcount.PopCountByPreAllocation8bitsLoop(v))
	fmt.Printf("popcount.PopCountByShiftingRight(%x) = %d\n", v, popcount.PopCountByShiftingRight(v))
	fmt.Printf("popcount.PopCountByShiftingLeft(%x) = %d\n", v, popcount.PopCountByShiftingLeft(v))
	fmt.Printf("popcount.PopCountByClearing(%x) = %d\n", v, popcount.PopCountByClearing(v))
	fmt.Printf("popcount.BitCount	(%x) = %d\n", v, popcount.BitCount(v))

}
