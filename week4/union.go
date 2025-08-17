package main

import "os"

func union(s1,s2 string)string{
	seen:=make(map[rune]bool)
	result:=""
	for _,char:=range s1+s2{
		if !seen[char]{
			seen[char]=true
			result+=string(char)
		}
	}
	return result
}

func main(){
	if len(os.Args)<3{
		os.Stdout.Write([]byte("\n"))
	}
	input:=os.Args[1]
	input2:=os.Args[2]
	result:=union(input,input2)
	os.Stdout.Write([]byte(result+"\n"))
}