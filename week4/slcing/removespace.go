package main

import "fmt"

func Removespace(s string)string{
	results:=""
	for i,char:=range s{
		if s[i]!=' '{
			results+=string(char)
		}
	}
	return results
}

func main(){
	fmt.Println(Removespace("hello world"))
}