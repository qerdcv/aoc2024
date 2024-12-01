package day2

import (
	"testing"
)

const testLine = "Game 100: 3 blue, 3 red, 6 green; " +
	"7 red, 2 green, 14 blue; " +
	"13 green, 9 red, 9 blue; " +
	"8 red, 10 green, 9 blue; " +
	"6 blue, 11 red; " +
	"3 blue, 3 red, 6 green; " +
	"7 red, 2 green, 14 blue; " +
	"13 green, 9 red, 9 blue; " +
	"8 red, 10 green, 9 blue; " +
	"6 blue, 11 red; " +
	"7 red, 2 green, 14 blue; " +
	"13 green, 9 red, 9 blue; " +
	"8 red, 10 green, 9 blue; " +
	"6 blue, 11 red; " +
	"3 blue, 3 red, 6 green; " +
	"7 red, 2 green, 14 blue; " +
	"13 green, 9 red, 9 blue; " +
	"8 red, 10 green, 9 blue; " +
	"6 blue, 11 red; " +
	"7 red, 2 green, 14 blue; " +
	"13 green, 9 red, 9 blue; " +
	"8 red, 10 green, 9 blue; " +
	"6 blue, 11 red; " +
	"3 blue, 3 red, 6 green; " +
	"7 red, 2 green, 14 blue; " +
	"13 green, 9 red, 9 blue; " +
	"8 red, 10 green, 9 blue; " +
	"6 blue, 11 red"

func BenchmarkParseValidGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseValidGame(testLine)
	}
}

func BenchmarkParseGamePower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseGamePower(testLine)
	}
}
