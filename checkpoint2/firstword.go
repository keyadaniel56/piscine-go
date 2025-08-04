package main

import "fmt"

func trim(s string)string{
	start:=0
	end:=len(s)-1

	for start<=end && s[start]==' '{
		start++
	}
	for end>=start && s[end]==' '{
		end--
	}
	return s[start:end+1]
}

func firstword(s string)string{
	word:=trim(s)
	start:=0
	end:=len(s)-1
	for start<=end && word[start]!=' '{
		start++
	}
	return s[:start]
}

func main(){
	fmt.Println(firstword("Hello world"))
}