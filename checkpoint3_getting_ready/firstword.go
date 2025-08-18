package main

import "fmt"

func trim_fields(s string) string {
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

func first_word(s string) string {
	words := trim_fields(s)
	start := 0
	for start >= 0 && words[start] != ' ' {
		start++
	}
	return words[:start]
}

func main() {
	fmt.Println(first_word("hello world i love coding"))
}
