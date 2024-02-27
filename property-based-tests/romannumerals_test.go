package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("1 is converted to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("2 is converted to II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
