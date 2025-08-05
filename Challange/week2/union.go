package main

import "fmt"

func union(s1 , s2 string)string{
	seen:=make(map[rune]bool)
	result:=""
	 
	 for _,char:=range s1+s2{
		if !seen[char]{
			result+=string(char)
			seen[char]=true
		}
	 }
	 return result
}


func main(){
	fmt.Println(union("padinton","paqefwtdjetyiytjneytjoeyjnejeyj"))
}