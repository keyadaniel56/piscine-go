package main

import "fmt"


func rstr(s string)string{
	runes:=[]rune(s)
	for i:=len(runes)-1;i>=0;i--{
		if (i==len(runes)-1 || s[i+1]==' ')  && runes[i]>='a' && runes[i]<='z'{
			runes[i]=runes[i]-32
		}
	}
	return string(runes)
}

func main(){
	fmt.Println(rstr("hello world i love coding"))
}