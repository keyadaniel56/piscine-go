package main

import "fmt"

func main() {
    fmt.Println(AddFront("Hello", []string{"world"}))
    fmt.Println(AddFront("Hello", []string{"world", "!"}))
    fmt.Println(AddFront("Hello", []string{}))
}



func AddFront(s string, slice []string) []string {
    // your code here
	if len(s)==0{
		return slice
	}
	return append([]string{s},slice...)
}

