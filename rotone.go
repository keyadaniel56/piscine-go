package main

import "fmt"

func rotone(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			shifted := (char-'a'+1)%26 + 'a'
			result += string(shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := (char-'A'+1)%26 + 'A'
			result += string(shifted)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(rotone("abcz"))
}
