package main

import "fmt"

func rotone(s string)string{
	b:=[]rune(s)
	for i,char:=range b{
		if char>='a' && char<='z'{
			b[i]=(char-'a'+1)%26+'a'
		}else if char>='A' && char<='Z'{
			b[i]=(char-'A'+1)%26+'A'
		}
	}
	return string(b)
}

func main(){
	word:="abc"
	fmt.Println(rotone(word))
}