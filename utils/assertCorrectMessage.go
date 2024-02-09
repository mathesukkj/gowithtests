package utils

import "testing"

func AssertCorrectMessage[T string | int](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
