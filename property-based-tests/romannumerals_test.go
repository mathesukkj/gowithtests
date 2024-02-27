package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 converted to I", 1, "I"},
		{"2 converted to II", 2, "II"},
		{"3 converted to III", 3, "III"},
		{"4 converted to IV", 4, "IV"},
		{"5 converted to V", 5, "V"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}
