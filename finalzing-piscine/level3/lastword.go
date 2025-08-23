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
	return s[start:end+1]
}

func lastword(s string) string {
	word := trimFields(s)
	end := len(word) - 1
	for end >= 0 && word[end] != ' ' {
		end--
	}
	return word[end+1:]
}

func main() {
	fmt.Println(lastword("hello world i love coding"))
}
