package main

import "fmt"


func ToUpper(s string) string {
	runes:=[]rune(s)
	for i,char:=range runes{
		if char>='a' && char<='z'{
			runes[i]+='A'-'a'
		}
	}
	return string(runes)
}



func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}