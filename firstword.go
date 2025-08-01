package main

import "os"

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

func firstword(s string)string{
	word:=TrimFields(s)
	start:=0
	for start>=0 && word[start]!=' '{
		start++
	}
	return word[:start]
}



func main(){
	if len(os.Args)<2{
		os.Stdout.Write([]byte("\n"))
	}
	result:=os.Args[1]
	output:=firstword(result)
	os.Stdout.Write([]byte(output + "\n"))
	
}