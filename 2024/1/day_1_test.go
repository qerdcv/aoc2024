package main

import (
	"strings"
	"testing"
)

// BenchmarkResolvePartOne-12    	  374170	      3112 ns/op
func BenchmarkResolvePartOne(b *testing.B) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	for range b.N {
		solvePartOne(strings.NewReader(input))
	}
}
