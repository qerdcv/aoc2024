package day3

import "testing"

var testData = [][]byte{
	[]byte("467..114.."),
	[]byte("...*......"),
	[]byte("..35..633."),
	[]byte("......#..."),
	[]byte("617*......"),
	[]byte(".....+.58."),
	[]byte("..592....."),
	[]byte("......755."),
	[]byte("...$.*...."),
	[]byte(".664.598.."),
}

func BenchmarkSearchParts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchParts(testData)
	}
}

func BenchmarkGearsTotalRatio(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchGearsTotalRatio(testData)
	}
}
