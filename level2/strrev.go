package main

import "fmt"

func strrev(s string) string {
	b := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
