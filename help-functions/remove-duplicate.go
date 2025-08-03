package main

import "fmt"

func remove_duplicates(s string) string {
	seen := make(map[rune]bool)
	result := ""
	for _, char := range s {
		if !seen[char] {
			seen[char] = true
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(remove_duplicates("abcdabcd"))
}
