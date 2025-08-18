package main

import "fmt"

func splitit(s string) []string {
	var words []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func sortWord(s []string) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
func longestWord(s string) string {
	words := splitit(s)
	if len(words) == 0 {
		return "\n"
	}
	sortWord(words)
	longest := words[0]
	for i := 1; i < len(words); i++ {
		if len(words[i]) > len(longest) {
			longest = words[i]
		}
	}
	return longest
}

func main() {
	fmt.Println(longestWord("hello world i love coding"))
}
