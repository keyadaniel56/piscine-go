package main

import "fmt"

func inter(s1, s2 string) string {
	result := ""
	seen := make(map[rune]bool)
	ins2 := make(map[rune]bool)

	// mark all characters that existd in s2
	for _, char := range s2 {
		ins2[char] = true
	}
	for _, char := range s1 {
		if ins2[char] && !seen[char] {
			result += string(char)
			seen[char] = true
		}
	}

	return result
}

func main() {
	fmt.Println(inter("abc", "fs9ruffcbas"))
}
