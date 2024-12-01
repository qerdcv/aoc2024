package days

import (
	"testing"
)

func BenchmarkDayOnePartOne(b *testing.B) {
	testData := []rune("543cshpxrfnnhonetkbhxtmlgczdndqjscb2mpftseven44five8nineeightrmgrljrljb8hxxfmdpbbvmblltxfive6mdsmm7mmpknprsix6five4vkqrsixmjkjps")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseOnlyDigits(testData)
	}
}

func BenchmarkDayOnePartTwo(b *testing.B) {
	testData := []rune("543cshpxrfnnhonetkbhxtmlgczdndqjscb2mpftseven44five8nineeightrmgrljrljb8hxxfmdpbbvmblltxfive6mdsmm7mmpknprsix6five4vkqrsixmjkjps")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseAlphaDigits(testData)
	}
}
