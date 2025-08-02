package main

import "fmt"

func capitalize(s string)string{
	runes:=[]rune(s)
	for i:=0;i<len(s);i++{
		if(i==0 || runes[i-1]==' ') &&(runes[i]>='a' && runes[i]<='z'){
			runes[i]=runes[i]-32
		}
	}
	return string(runes)
}

func main(){
	result:="hello world i love coding so much"
	fmt.Println(capitalize(result))
}