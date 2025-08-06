package main

import "fmt"

func TrimFields(s string)string{
	start:=0
	end:=len(s)-1
	for start<=end && s[start]==' '{
		start ++
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
			words=append(words,word)
			word=""
		}
	}
	if word!=""{
		words=append(words,word)
	}
	return words
}


func separator(s string, separator string)string{
	words:=split(s)
	result:=""
	for i:=0;i<len(words);i++{
		result+=words[i]
		if i!=len(words)-1{
			result+=separator
		}
	}
	return result
}

func expandstr(s string)string{
	result:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' ' && s[i]!='\t' && s[i]!='\n'{
                result+=string(s[i])
		}else{
			if len(result)>0 && result[len(result)-1]!=' '{
				result+=" "
			}
		}
	}
	
	return result
}


func expand(s string)string{
	word:=expandstr(s)
	spaces:="   "
	result:=separator(word,spaces)
	return result
}
func main(){
	fmt.Println(split("hello world"))
	fmt.Println(separator("hello world","_"))
	fmt.Println(expandstr("           hello      world     "))
		fmt.Println(expand("           hello      world    i    love       coding "))
}