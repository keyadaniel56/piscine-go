package main

import "fmt"

func split_string(s string)[]string{
	var words [] string
	word:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			word+=string(s[i])
		}else if word!=" "{
			words=append(words,word)
			word=" "
		}
	}
	if word!=""{
		words=append(words,word)
	}
	return words
}

func main(){
	fmt.Println(split_string("hello world"))
}