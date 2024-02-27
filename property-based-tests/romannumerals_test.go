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
		{"9 converted to IX", 9, "IX"},
		{"10 converted to X", 10, "X"},
		{"14 converted to XIV", 14, "XIV"},
		{"18 converted to XVIII", 18, "XVIII"},
		{"20 converted to XX", 20, "XX"},
		{"39 converted to XXXIX", 39, "XXXIX"},
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
