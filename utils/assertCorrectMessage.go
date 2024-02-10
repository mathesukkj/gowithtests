package utils

import (
	"reflect"
	"testing"
)

func AssertCorrectMessageNoParams[T string | int | float64](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertCorrectMessageWithParam[T string | int, P []int](t testing.TB, got, want T, param P) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v, %v given", got, want, param)
	}
}

func AssertCorrectMessageSlice(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didnt get one")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}
