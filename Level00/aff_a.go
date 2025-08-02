package main

import "github.com/01-edu/z01"

func main() {
	str := "hello"
	if len(str) == 0 {
		z01.PrintRune('a')
	}
	for _, char := range str {
		if char == 'a' {
			z01.PrintRune('a')
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('\n')
		}
	}
}
