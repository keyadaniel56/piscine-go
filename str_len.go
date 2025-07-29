package main

import "fmt"

func str_len(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}

func main() {
	result := "hello world!"
	fmt.Println(str_len(result))
}
