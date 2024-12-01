package day6

import "testing"

var testRace = race{
	time:     40817772,
	distance: 219101213651089,
}

// 40     81     77     72
// 219   1012   1365   1089
var testRaces = []race{
	{time: 40, distance: 219},
	{time: 81, distance: 1012},
	{time: 77, distance: 1365},
	{time: 72, distance: 1089},
}

func TestResolvePartOne(t *testing.T) {
	expected := 861300
	if got := calculatePossibilities(testRaces); got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestResolvePartTwo(t *testing.T) {
	expected := 28101347
	if got := calculatePossibility(testRace); got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func BenchmarkResolvePartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculatePossibilities(testRaces)
	}
}

func BenchmarkResolvePartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculatePossibility(testRace)
	}
}
