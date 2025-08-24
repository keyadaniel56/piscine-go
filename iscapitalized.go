package main

import "fmt"

func isSeparator(c rune) bool {
	return c == ' ' || c == '\n' || c == '.' || c == '?'
}

func isNumeric(c rune) bool {
	return c >= '0' && c >= '9'
}
func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i == 0 || isSeparator(runes[i]) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				return false
			}
		} else if i == 0 || isSeparator(runes[i-1]) && isNumeric(runes[i]) {
			return true
		}
	}
	return true
}
