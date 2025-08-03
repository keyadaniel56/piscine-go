package main

import "fmt"

func contains_char(s string, c rune)bool{
	for _,char:=range s{
		if char==c{
			return true
		}
	}
	return false
}

func inter(s1,s2 string)string{
	result:=""
	for _,char:=range s1{
		if contains_char(s2 , char) && !contains_char(result,char){
			result+=string(char)
		}
	}
	return result
}

func main(){
	fmt.Println(inter("hello","worlld"))
}