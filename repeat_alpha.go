package main

import "fmt"

func repeat_alpha(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count := int(char - 'a' + 1)
			for i := 0; i < count; i++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			count := int(char - 'A' + 1)
			for i := 0; i < count; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	result := "Abc"
	fmt.Println(repeat_alpha(result))
}
