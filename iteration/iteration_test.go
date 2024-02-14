package main

import (
	"fmt"
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaaa"

	utils.AssertCorrectMessageNoParams(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 20)
	fmt.Println(result)
	// Output: aaaaaaaaaaaaaaaaaaaaa
}
