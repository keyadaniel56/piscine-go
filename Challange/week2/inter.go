package main

import "fmt"


func inter(s,s2 string)string{
	set1:=make(map[rune]bool)
	seen:=make(map[rune]bool)
    result:=""
	for _,char:=range s2{
		set1[char]=true
	}

	for _,char:=range s{
		if set1[char] && !seen[char]{
			result+=string(char)
			seen[char]=true
		}
	}
	return result
}
func main(){
	fmt.Println(inter("padinton","paqefwtdjetyiytjneytjoeyjnejeyj"))
}