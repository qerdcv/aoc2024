package day4

import "testing"

var testData = []numbersPair{
	{
		cardNumbers: map[int]bool{
			41: true, 48: true, 83: true, 86: true, 17: true,
		},
		winningNumbers: map[int]bool{
			83: true, 86: true, 6: true, 31: true, 17: true, 9: true, 48: true, 53: true,
		},
	},
	{
		cardNumbers: map[int]bool{
			13: true, 32: true, 20: true, 16: true, 61: true,
		},
		winningNumbers: map[int]bool{
			61: true, 30: true, 68: true, 82: true, 17: true, 32: true, 24: true, 19: true,
		},
	},
}

func BenchmarkResolvePartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getTotalCardValues(testData)
	}
}

func BenchmarkResolvePartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getTotalCardCopies(testData)
	}
}
