package iteration

import (
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaaa"

	utils.AssertCorrectMessage(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
