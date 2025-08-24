package main

import "fmt"

func isCamel(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		char := s[i]
		// none digit character not alowd also digit
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
		//uppercase cannot follow one another
		if (i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z') && char >= 'A' && char <= 'Z' {
			return false
		}
		//cannot end with uppercase
		if i == len(s)-1 && char >= 'A' && char <= 'Z' {
			return false
		}
	}
	return true
}

func camelToSnakeCae(s string) string {
	if !isCamel(s) {
		return s
	}
	result := []rune{}
	for i, char := range s {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result = append(result, '_')
			result = append(result, char)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
func main() {
	fmt.Println(camelToSnakeCae("helloWorld"))
}
