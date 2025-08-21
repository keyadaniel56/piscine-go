package main

import (
	"fmt"

	
)

func removeTrailingSpace(s string)string{
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

func join(words []string)string{
	results:=""
	for i:=0;i<len(words);i++{
		if i>0{
			results+=" "
		}
		results+=words[i]
	}
	return results
}


func firstword(s string)string{
	words:=removeTrailingSpace(s)
	start:=0
	for start>=0 && words[start]!=' '{
		start++
	}
	return words[:start]
}


func lastword(s string)string{
	words:=removeTrailingSpace(s)
	end:=len(words)-1
	for end>=0 && words[end]!=' '{
		end--
	}
	return words[end+1:]
}

func flipfirstword(s string)string{
	words:=split(s)
	if len(words)==0{
		return ""
	}
	first:=words[0]
	rest:=words[1:]
	flipped:=append(rest,first)
	return join(flipped)
}

func fliplastword(s string)string{
	words:=split(s)
	if len(words)==0{
		return ""
	}
	last:=words[len(words)-1]
	rest:=words[:len(words)-1]
	flipped:=append([]string{last},rest...)
	return join(flipped)
}
func main(){
	// fmt.Println(firstword("hello i love coding"))
	// fmt.Println(lastword("hello i love coding"))
	// fmt.Println(split("hello world i love coding"))
	// fmt.Println(flipfirstword("hello world i love coding"))
	fmt.Println(fliplastword("hello world i love coding"))
}