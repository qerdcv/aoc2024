package days_test

import (
	"strings"
	"testing"

	day "github.com/qerdcv/aoc/days/day_1"
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
		day.ResolvePartOne(strings.NewReader(input))
	}
}
