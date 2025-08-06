package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	
	var words [] string
	word:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			word+=string(s[i])
		}else if word!=""{
			words=append(words,word)
			word=""
		}
	}
	if word!=""{
		words=append(words,word)
	}
	return words
}


func separate(s , separator string)string{
	result:=""
	word:=SplitWhiteSpaces(s)
	for i:=0;i<len(word);i++{
		result+=word[i]
		if i!=len(word)-1{
			result+=separator
		}
	}
	return result
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you? i love coding so mch"))
}