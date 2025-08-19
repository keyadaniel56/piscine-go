package main

import "fmt"

func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G\n"
	}
	if len(str) == 0 {
		return "\n"
	}
	return "Invalid Input" + "\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
