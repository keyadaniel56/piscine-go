package main

import "fmt"

func RepeatAlpha(s string) string {
	if len(s) == 0 {
		return s
	}
	results := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' || (char >= 0 && char <= 9) {
			count := int(char - 'a' + 1)
			for i := 0; i < count; i++ {
				results += string(char)
			}
		} else if char >= 'A' && char <= 'Z' || (char >= 0 && char <= 9) {
			count := int(char - 'A' + 1)
			for i := 0; i < count; i++ {
				results += string(char)
			}
		} else {
			results += string(char)
		}
	}
	return results
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
