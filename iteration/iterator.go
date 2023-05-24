package iteration

import "testing"

var repeat int = 5

func Repeat(character string, repeatInput int) string {
	var repeated string
	if repeatInput >= 0 {
		repeat = repeatInput
	}
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
