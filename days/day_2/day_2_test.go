package day2

import "testing"

// BenchmarkResolveIsSaveV1-12    	1000000000	         1.192 ns/op	       0 B/op	       0 allocs/op
func BenchmarkResolveIsSaveV1(b *testing.B) {
	input := []int{80, 80, 79, 73, 71, 67}

	for range b.N {
		isSave(input)
	}
}

// BenchmarkResolveIsSaveV2-12    	20447596	        57.04 ns/op	      48 B/op	       1 allocs/op
func BenchmarkResolveIsSaveV2(b *testing.B) {
	input := []int{80, 80, 79, 73, 71, 67}

	for range b.N {
		isSaveV2(input)
	}
}
