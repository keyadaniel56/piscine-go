package main

import "fmt"

func isnumeric(n rune) bool {
	return n >= '0' && n <= '9'
}

func isseparator(c rune) bool {
	return c == ' ' || c == '\n' || c == '!'
}

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i == 0 || isseparator(runes[i-1]) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				return false
			}
		} else if i == 0 || isseparator(runes[i-1]) && isnumeric(runes[i-1]) {
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
