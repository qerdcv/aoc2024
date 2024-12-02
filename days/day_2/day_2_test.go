package day2

import "testing"

func BenchmarkResolveIsSaveV1(b *testing.B) {
	input := []int{80, 80, 79, 73, 71, 67}

	for range b.N {
		isSave(input)
	}
}

func BenchmarkResolveIsSaveV2(b *testing.B) {
	input := []int{80, 80, 79, 73, 71, 67}

	for range b.N {
		isSave(input)
	}
}
