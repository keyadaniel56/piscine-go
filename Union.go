package main

import "fmt"

func union(str1, str2 string) string {
	seen := make(map[rune]bool)
	result := ""
	for _, char := range str1 + str2 {
		if !seen[char] {
			result += string(char)
			seen[char] = true
		}
	}
	return result
}

func main() {
	fmt.Println(union("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
}
