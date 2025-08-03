package main

import "fmt"

func split_strings(s string) []string {
	var words []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else if word != " " {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func separate(s, separator string) string {
	word := split_strings(s)
	result := ""
	for i := 0; i < len(word); i++ {
		result += word[i]
		if i != len(word)-1 {
			result += separator
		}
	}
	return result
}

//for testing
func main() {
	fmt.Println(separate(" hello world", "_"))
}
