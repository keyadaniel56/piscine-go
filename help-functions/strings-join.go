package main

import "fmt"

func split_str(s string)[] string{
 var words []string
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

func join_string(s , sepa string)string{
	word:=split_str(s)
	result:=""
	for i:=0;i<len(word);i++{
		result+=word[i]
		if i!=len(word)-1{
			result+=sepa
		}
	}
	return result
}
func main(){
  fmt.Println(join_string("hello world"," "))
}