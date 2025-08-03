package main

import "fmt"
func capitalize_last(s string)string{
	runes:=[]rune(s)
	for i:=0;i<len(runes);i++{
		if runes[i]>='a' && runes[i]<='z'{
			if i+1==len(runes) || runes[i+1]==' '{
				runes[i]=runes[i]-32
			}
		}
	}
	return string(runes)
}

// testing

func main(){
	fmt.Println(capitalize_last("hello world i love coding"))
}