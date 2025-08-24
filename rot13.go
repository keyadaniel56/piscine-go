package main

import "fmt"

func rot13(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(char + 13)
		} else if char >= 'A' && char <= 'Z' {
			result += string(char + 13)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(rot13("abAc"))
}
