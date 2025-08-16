package main

import "fmt"

func trimeFields(s string) string {
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

func LastWord(s string) string {
	word := trimeFields(s)
	start := 0
	end := len(word) - 1
	for end >= start && word[end] != ' ' {
		end--
	}
	return word[end+1:] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
	fmt.Print(LastWord(" i love coding i am cdk"))
}
