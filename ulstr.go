package main

import "fmt"

func ulstr(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32)
		} else if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(ulstr("abcD"))
}
