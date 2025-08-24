package main

import "fmt"

func alphamirror(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			mirrored := 'z' - (char - 'a')
			result += string(mirrored)
		} else if char >= 'A' && char <= 'Z' {
			mirrored := 'Z' - (char - 'A')
			result += string(mirrored)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(alphamirror("zZaA"))
}
