package main

import "fmt"

func ToLower(s string)string{
	runes:=[]rune(s)
	for i,char:=range runes{
		if char>='A' && char<='Z'{
			runes[i]+='a'-'A'
		}
	}
	return string(runes)
}


func main() {
	fmt.Println(ToLower("Hello! How are you? do you love coding"))
}