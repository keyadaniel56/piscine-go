package main

import "fmt"

func StringToBool(s string) bool {
	if s == "True" || s == "T" || s == "t" {
		return true
	}
	return false
}

func main() {
	fmt.Println(StringToBool("True"))
	fmt.Println(StringToBool("T"))
	fmt.Println(StringToBool("False"))
	fmt.Println(StringToBool("TTFF"))
}
