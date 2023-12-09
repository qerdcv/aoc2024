package day9

import "testing"

// 0 3 6 9 12 15
// 1 3 6 10 15 21
// 10 13 16 21 30 45
var testHistories = [][]int{
	{0, 3, 6, 9, 12, 15},
	{1, 3, 6, 10, 15, 21},
	{10, 13, 16, 21, 30, 45},
}

func BenchmarkResolvePartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		extrapolateHistories(testHistories, forwardExtrapolator)
	}
}

func BenchmarkResolvePartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		extrapolateHistories(testHistories, backwardExtrapolator)
	}
}
