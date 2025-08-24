package main

import "fmt"

func CleanSpace(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == ' ' && s[i-1] == ' ' {
			continue
		}
		result += string(s[i])
	}
	return result
}

func remove_trail(s string) string {
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

func expandstr(s string) string {
	word := remove_trail(s)
	cleaned := CleanSpace(word)
	return cleaned
}
func main() {
	fmt.Println(expandstr("hello   world"))
}
