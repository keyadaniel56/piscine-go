package main

import (
	//"fmt"
	"os"
)

func ts(s string)string{
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


func rs(s string)string{
	b:=[]rune(s)
	result:=""
	for i:=0;i<len(b);i++{
		if i>0 && b[i]==' ' && b[i-1]==' '{
			continue
		}
		result+=string(b[i])
	}
	return result
}


func sr(s ,sep string)string{
	b:=[]rune(s)
	result:=""
	for i:=0;i<len(b);i++{
		if i>0 && b[i]==' ' && b[i-1]==' '{
			continue
		}
		if b[i]==' '{
			result+=sep
		}else{
			result+=string(b[i])
		}
	}
	return result
}


func cleanstr(s string)string{
	words:=ts(s)
	cleaned:=rs(words)
	return cleaned
}


func main(){
	if len(os.Args)<2{
		os.Stdout.Write([]byte("\n"))
		return
	}
	input:=os.Args[1]
	results:=cleanstr(input)
	os.Stdout.Write([]byte(results))
	os.Stdout.Write([]byte("\n"))
}