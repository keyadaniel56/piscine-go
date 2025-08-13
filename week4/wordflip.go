package main

import "fmt"


func split(s string)[]string{
	var words [] string
	word:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			word+=string(s[i])
		}else if word!=""{
			words=append(words, word)
			word=""
		}
	}
	if word!=""{
		words=append(words, word)
	}
	return words
}

func join(words [] string)string{
	result:=""
	for _,word:=range words{
		result+=word + " "
	}
	return result
}


func WordFlip(str string) string {
	words:=split(str)
	result:=""
	for i,j:=0,len(words)-1;i<j;i,j=i+1,j-1{
		words[i],words[j]=words[j],words[i]
	}
	result=join(words)
	return result+"\n"
}
func main(){
	fmt.Println(WordFlip("hello world i love coding"))
}