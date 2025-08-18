package main

import "fmt"

func split(s string)[]string{
	var words [] string
	word:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			word+=string(s[i])
		}else if word!=""{
			words = append(words, word)
			word=""
		}
	}
	if word!=""{
		words=append(words, word)
	}
	return words
}


func countWords(s string)int{
	words:=split(s)
	count:=0;
	for i:=0;i<len(words);i++{
		count++
	}
	return count
}

func main(){
	fmt.Println(countWords("hello world i love coding"))
}