package main

import "fmt"

func trimFields(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && s[start] == ' ' {
		start++
	}
	for end >= start && s[end] == ' ' {
		end--
	}
	return s[start : end+1]
}

func firstWord(s string) string {
	word := trimFields(s)
	start := 0
	for start >= 0 && word[start] != ' ' {
		start++
	}
	return word[:start]
}
func main() {
	fmt.Println(firstWord("    hello world    "))
}
