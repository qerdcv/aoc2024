package day3

import "testing"

// BenchmarkResolvePartOne-12    	14404603	        85.16 ns/op	       0 B/op	       0 allocs/op
func BenchmarkResolvePartOne(b *testing.B) {
	input := []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)
	for range b.N {
		ResolvePartOne(input)
	}
}

// BenchmarkResolvePartTwo-12    	12893824	        94.75 ns/op	       0 B/op	       0 allocs/op
func BenchmarkResolvePartTwo(b *testing.B) {
	input := []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)
	for range b.N {
		ResolvePartTwo(input)
	}
}
