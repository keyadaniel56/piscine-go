package main

import "fmt"

func isSeparator(c rune) bool {
	return c == ' ' || c == '\n'
}

func isNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i == 0 || isSeparator(runes[i-1]) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				return false
			}
		} else if i == 0 || isSeparator(rune(i-1)) && isNumeric(runes[i]) {
			return true
		}
	}
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
