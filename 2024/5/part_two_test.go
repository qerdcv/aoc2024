package main

import "testing"

func BenchmarkFixUpdate(b *testing.B) {
	rules := map[int][]int{
		53: {47, 75, 61, 97},
		13: {97, 61, 27, 47, 75, 53},
		61: {91, 47, 75},
		47: {97, 75},
		29: {75, 97, 53, 61, 47},
		75: {97},
	}
	update := []int{97, 13, 75, 29, 47}

	for range b.N {
		fixUpdate(rules, update)
	}
}
