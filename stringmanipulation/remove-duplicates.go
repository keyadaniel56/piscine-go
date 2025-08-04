package main 

import "fmt"

func removeduplicates(s string)string{
	seen:=make(map[rune]bool)
	result:=""
	for _,char:=range s{
		if !seen[char]{
			result+=string(char)
			seen[char]=true
		}
	}
	return result
}

func main(){
	fmt.Println(removeduplicates("aabccde"))
}