package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
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

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
