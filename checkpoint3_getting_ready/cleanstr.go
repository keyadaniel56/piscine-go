package main

import "fmt"

func trimfields(s string) string {
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

func rmv_space(s string) string {
	runes := []rune(s)
	result := ""
	for i := 0; i < len(runes); i++ {
		if i > 0 && runes[i-1] == ' ' && runes[i] == ' ' {
			continue
		}
		result += string(runes[i])
	}
	return result
}

func cleanstr(s string) string {
	word := trimfields(s)
	cleaned := rmv_space(word)
	return cleaned
}
func main() {
	fmt.Println(cleanstr("hello     world     i   love      coding"))
}
