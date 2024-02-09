package integers

import (
	"mathesukkj/gowithtests/utils"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	utils.AssertCorrectMessage(t, got, want)
}
