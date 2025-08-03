package main

import "fmt"


func char_contain(s string, c rune)bool{
	for _,char:=range s{
		if char==c{
			return true
		}
	}
	return false
}


func wrdmtch(s1,s2 string)string{
	result:=""
	for _,char:=range s1{
		if char_contain(s2,char) && !char_contain(result,char){
			result+=string(char)
		}
	}
	return result
}


func main(){
	fmt.Println(wrdmtch("faya","fgvvfdxcacpolhyghbred"))
}