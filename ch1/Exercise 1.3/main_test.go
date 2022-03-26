package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	args := [...]string{"these", "are", "my", "5", "arguments"};
	for i := 0; i < b.N; i++ {
		echo1(args[:]);
	}
}

func BenchmarkEcho2(b *testing.B) {
	args := [...]string{"these", "are", "my", "5", "arguments"};
	for i := 0; i < b.N; i++ {
		echo2(args[:]);
	}
}

func BenchmarkEcho3(b *testing.B) {
	args := [...]string{"these", "are", "my", "5", "arguments"};
	for i := 0; i < b.N; i++ {
		echo3(args[:]);
	}
}
