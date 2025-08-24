package main

import "fmt"

func space_clean(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == ' ' && s[i-1] == ' ' {
			continue
		}
		result += string(s[i])
	}
	return result
}

func Expandstr(s string) string {
	spaces := "   "
	result := ""
	words := space_clean(s)
	for i := 0; i < len(words); i++ {
		if words[i] == ' ' {
			result += spaces
		} else {
			result += string(words[i])
		}
	}
	return result
}

func main() {
	fmt.Println(Expandstr("hello world i love coding"))
}
