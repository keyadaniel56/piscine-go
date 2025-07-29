package main

import "os"

func rev_str(s string) string {
	b := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	input := os.Args[1]
	result := rev_str(input)
	os.Stdout.Write([]byte(result + "\n"))
}
