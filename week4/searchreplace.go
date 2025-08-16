package main

import (
	"os"
)

func searchreplace(input, search, replace string) string {
	inputRune := []rune(input)
	searchRune := []rune(search)
	replaceRune := []rune(replace)

	if len(searchRune) != 1 || len(inputRune) != 1 {
		return input
	}
	s := searchRune[0]
	r := replaceRune[0]
	result := ""
	for _, char := range inputRune {
		if char == s {
			result += string(r)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	if len(os.Args) < 4 {
		os.Stdout.Write([]byte("\n"))
	}
	input := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]

	result := searchreplace(input, search, replace)
	os.Stdout.Write([]byte(result + "\n"))
}
