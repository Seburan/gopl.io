// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"gopl.io/ch2/exercises/popcount"
)

// -- Benchmarks --

func BenchmarkPopCountByPreAllocation8bits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByPreAllocation8bits(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByPreAllocation8bitsLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByPreAllocation8bitsLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShiftingLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShiftingLeft(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShiftingRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByShiftingRight(0x1234567890ABCDEF)
	}
}
func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClearing(0x1234567890ABCDEF)
	}
}



// Go 1.6, 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// Go 1.6, 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// Go 1.7, 3.5GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-12                 2000000000        0.28 ns/op
// BenchmarkBitCount-12                 2000000000        0.27 ns/op
// BenchmarkPopCountByClearing-12       100000000        18.5 ns/op
// BenchmarkPopCountByShifting-12       20000000         70.1 ns/op
