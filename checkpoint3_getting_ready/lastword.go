package main

import "fmt"

func trim(s string) string {
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

func lastword(s string) string {
	words := trim(s)
	end := len(words) - 1
	for end >= 0 && words[end] != ' ' {
		end--
	}
	return words[end+1:] + "\n"
}

func main() {
	fmt.Print(lastword("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(lastword(" lorem,ipsum "))
	fmt.Print(lastword(" "))
}
