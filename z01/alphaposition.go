package main

import "fmt"

func AlphaPosition(c rune) int {
	// your code goes here
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 1)
	}
	return 0
}

func main() {
	fmt.Println(AlphaPosition('b'))
}
