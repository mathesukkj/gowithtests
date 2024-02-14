package main

import (
	"fmt"
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	utils.AssertCorrectMessageNoParams(t, got, want)
}

func ExampleAdd() {
	sum := Add(10, 10)
	fmt.Println(sum)
	// Output: 20
}
