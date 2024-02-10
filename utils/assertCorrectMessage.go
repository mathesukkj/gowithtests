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
