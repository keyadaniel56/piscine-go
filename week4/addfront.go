package main

import "fmt"

func AddFront(s string, slice []string) []string {
	// your code here
	return append([]string{s}, slice...)
}

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}
