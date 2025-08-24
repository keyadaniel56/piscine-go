package main

import "fmt"

func inter(str1, str2 string) string {
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)
	seen := make(map[rune]bool)
	result := ""

	for _, char := range str1 {
		set1[char] = true
	}
	for _, char := range str2 {
		set2[char] = true
	}

	for _, char := range str1 {
		if !seen[char] && set2[char] {
			result += string(char)
			seen[char] = true
		}
	}

	for _, char := range str2 {
		if !seen[char] && set1[char] {
			result += string(char)
			seen[char] = true
		}
	}
	return result
}

func main() {
	fmt.Println(inter("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
}
