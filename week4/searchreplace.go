package main

import (
	"os"
)

func searchreplace(input, search, replace string) string {
	result := ""
	if len(search) != 1 || len(replace) != 1 {
		return input
	}
	s := search[0]
	r := replace[0]
	for i := 0; i < len(input); i++ {
		if input[i] == s {
			result += string(r)
		} else {
			result += string(input[i])
		}
	}
	return result
}

func main() {
	if len(os.Args) < 4 {
		os.Stdout.Write([]byte("\n"))
	}
	input := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]
	os.Stdout.Write([]byte(searchreplace(input, search, replace) + "\n"))
}
