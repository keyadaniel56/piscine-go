package main 

import "fmt"

func Trim(s string)string{
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


func rmv_trailing(s string)string{
	result:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			result+=string(s[i])
		}else if i>0 && s[i-1]!=' '{
			result+=" "
		}
	}
	return result
}


func split(s string)[]string{
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
func separator(s ,separator string)string{
	result:=""
	words:=split(s)
	for i:=0;i<len(words);i++{
		result+=string(words[i])
		if i!=len(words)-1{
			result+=separator
		}
	}
	return result
}


func expand(s string)string{
	input:=rmv_trailing(s)
	spaces:="   "
	result:=separator(input,spaces)
	return result
}
func main(){
	//fmt.Println(rmv_trailing("     hello     world    i   love     coding      "))
	fmt.Println(separator("hello world","_"))
	fmt.Println(expand("hello        world"))
}