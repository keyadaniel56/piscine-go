package main

import "fmt"

func Remove_Trailing_Space(s string) string {
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

func firstword(s string) string {
	word := Remove_Trailing_Space(s)
	start := 0
	for start >= 0 && word[start] != ' ' {
		start++
	}
	return word[:start]
}

func main() {
	fmt.Println(firstword("hello world"))
}
