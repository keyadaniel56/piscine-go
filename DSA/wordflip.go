package main

import "fmt"

func Split(s string) []string {
	words := []string{}
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

func Join(words []string) string {
	results := ""
	for i := 0; i < len(words); i++ {
		results += words[i] + " "
	}
	if len(results) > 0 {
		results = results[:len(results)-1]
	}
	return results
}
func WordFlip(str string) string {
	word := Split(str)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return Join(word)
}

func main() {
	fmt.Println(WordFlip("hello i love coding"))
}
