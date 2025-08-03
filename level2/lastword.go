package main

import "fmt"

func splits(s string) string {
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

func lastwords(s string) string {
	word := splits(s)
	end := len(word) - 1
	for end >= 0 && word[end] != ' ' {
		end--
	}
	return s[end+1:]
}

func main() {
	fmt.Println(lastwords("hello world"))
}
