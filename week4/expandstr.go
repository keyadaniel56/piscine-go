package main

import (
	"fmt"
)

func trimFields(s string)string{
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


func trimspace(s string)string{
	b:=[]rune(s)
	result:=""
	for i:=0;i<len(s);i++{
		if i>0 && b[i-1]==' ' && b[i]==' '{
			continue
		}
		result+=string(b[i])
	}
	return result
}

func add_space(s ,sep string)string{
	b:=[]rune(s)
	result:=""
	for i:=0;i<len(s);i++{
		if i>0 && b[i-1]==' ' && b[i]==' '{
			continue
		}
		if b[i]==' '{
			result+=string(sep)
		}else{
			result+=string(b[i])
		}
	}
	return result
}


func expandstr(s string)string{
	words:=trimFields(s)
	trimed:=trimspace(words)
	extr_space:="   "
	results:=add_space(trimed,extr_space)
	return results

}

func main(){
	fmt.Println(expandstr("hello world i      love coding"))
}