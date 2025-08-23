package main

import "fmt"

func search_replace(input, search, replace string) string {
	//input_rune:=[]rune(input)
	replace_rune := []rune(replace)
	search_rune := []rune(search)

	if len(replace_rune) != 1 || len(replace_rune) != 1 {
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
	return string(result)
}

func main() {
	fmt.Println(search_replace("hello world", "h", "j"))
}
