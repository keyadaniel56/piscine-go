package main

import "fmt"

func capitalizer(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if (i == 0 || runes[i-1] == ' ') && (runes[i] >= 'a' && runes[i] <= 'z') {
			runes[i] -= 32
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(capitalizer("hello world i love coding"))
}
