package main

import "fmt"

func RemoveOdd(s string) string {
	// your code here
	result := ""
	for i, char := range s {
		if char == ' ' {
			result += string(char)
		} else if i%2 == 0 {
			result += string(char)
		}
	}
	return result
}
