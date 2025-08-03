package main

import "fmt"

func TrimFields(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && s[start] == ' ' {
		start++
	}
	for end >= start && s[end] == ' ' {
		end--
	}
	return s[start : end+1]
}
func main() {
	fmt.Println(TrimFields("   hello world    "))
}
