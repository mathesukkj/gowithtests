package main

import "strings"

func ConvertToRoman(num int) string {
	var results strings.Builder

	for range num {
		results.WriteString("I")
	}

	return results.String()
}
