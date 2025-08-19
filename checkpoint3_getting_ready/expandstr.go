package main

import "fmt"

func remove_trailing(s string)string{
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

func removespace(s string)string{
	runes:=[]rune(s)
	result:=""
	for i:=0;i<len(runes);i++{
		if i>0 && runes[i-1]==' ' && runes[i]==' '{
			continue
		}
		result+=string(runes[i])
	}
	return result
}


func add_space(s,sep string)string{
	runes:=[]rune(s)
	results:=""
	for i:=0;i<len(runes);i++{
		if i>0 && runes[i-1]==' ' && runes[i]==' '{
			continue
		}
		if s[i]==' '{
			results+=string(sep)
		}else{
			results+=string(s[i])
		}
	}
	return results
}
func expandstr(s string)string{
	words:=remove_trailing(s)
	trimmed:=removespace(words)
	spaces:="   "
	results:=add_space(trimmed,spaces)
	return results
}


func main(){
	fmt.Println(removespace("hello   world"))
	fmt.Println(expandstr("hello world i love coding"))
}