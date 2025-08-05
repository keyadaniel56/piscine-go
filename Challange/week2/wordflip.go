package main

import "fmt"


func TrimFields(s string)string{
	start:=0
	end:=len(s)-1

	for start<=end && s[start]==' '{
		start++
	}
	for end>=start && s[end]==' '{
		end--
	}
	return s[start:end+1]
}



func split(s string)[]string{
	var words [] string
	word:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			word+=string(s[i])
		}else if word !=""{
			words=append(words,word)
			word=""
		}
	}
	if word!=""{
		words=append(words,word)
	}
	return words
}

func flip_word(s string)string{
	word:=split(s)
	for i,j:=0,len(word)-1;i<j;i,j=i+1,j-1{
		word[i],word[j]=word[j],word[i]
	}
	output:=""
	for _,char:=range word{
		output+=string(char)
		output+=" "
	}
	final:=TrimFields(output)
	return final
}


func main(){
	word:="hello world i love coding"
	fmt.Println(split(word))
	fmt.Println(flip_word(word))
}