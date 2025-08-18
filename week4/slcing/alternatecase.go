package main

import "fmt"

func alternatercase(s string) string {
	result := ""
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			if i%2 == 0 {
				result += string(char - 32) // make uppercase
			} else {
				result += string(char) // keep lowercase
			}
		} else if char >= 'A' && char <= 'Z' {
			if i%2 == 1 {
				result += string(char + 32) // make lowercase
			} else {
				result += string(char) // keep uppercase
			}
		} else {
			result += string(char) // keep spaces/punctuation unchanged
		}
	}
	return result
}

func main() {
	fmt.Println(alternatercase("hello world"))
}
