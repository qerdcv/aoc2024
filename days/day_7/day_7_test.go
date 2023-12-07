package day7

import (
	"strings"
	"testing"
)

// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483
// [ {[13 10 0 0 10] 220 0} {[12 12 12 0 14] 483 0}]

var testData = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestResolvePartOne(t *testing.T) {
	hands := parseHands(strings.NewReader(testData), cardToValue)
	expected := 6440
	if got := calculateTotalBid(hands, false); got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestResolvePartTwo(t *testing.T) {
	hands := parseHands(strings.NewReader(testData), cardToValueP2)
	expected := 5905
	if got := calculateTotalBid(hands, true); got != expected {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func BenchmarkResolvePartOne(b *testing.B) {
	hands := parseHands(strings.NewReader(testData), cardToValue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateTotalBid(hands, false)
	}
}

func BenchmarkResolvePartTwo(b *testing.B) {
	hands := parseHands(strings.NewReader(testData), cardToValue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateTotalBid(hands, true)
	}
}
