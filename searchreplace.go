package main

import "fmt"

func searchReplace(input, search, replace string) string {
	search_rune := []rune(search)
	replace_rune := []rune(replace)

	if len(search_rune) != 1 || len(replace_rune) != 1 {
		return input
	}
	s := search_rune[0]
	r := replace_rune[0]
	result := ""

	for _, char := range input {
		if char == s {
			result += string(r)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(searchReplace("hello world", "h", "j"))
}
