package main

import "fmt"


func  str_capitalizer(s string)string{
	runes:=[]rune(s)
	for i:=0;i<len(runes);i++{
		if(i==0 || !isAlphanumeric(runes[i-1])) && (runes[i]>='a' && runes[i]<='z'){
			runes[i]=runes-32
		}
		return string(runes)
	}
}

func isAlphanumeric(r runes)bool{
	return(r>='A' && r<='Z')||
	(r>='z'&&r<='z')||
	(r>='0' && r<='9')
}