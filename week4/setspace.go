package main

import "fmt"


func capitalize(s string)string{
	runes:=[]rune(s)
	for i:=0;i<len(runes);i++{
		for (i==0 || runes[i-1]==' ') && (runes[i]>='a' && runes[i]<='z'){
			runes[i]-=32
		}
	}
	return string(runes)
}


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

func Trimspace(s string)string{
	result:=""
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			result+=string(s[i])
		}
	}
	return result
}

func SetSpace(s string) string {
	if len(s)==0{
		return "Error"
	}
	word:=capitalize(s)
	trimmed:=Trimspace(word)
    for _,char:=range trimmed{
		if char>='0' && char<='9' || char==' ' && char>='a' && char<='z'{
			return "Error"
		}
	}

	var result string
	for i:=0;i<len(trimmed);i++{
		if i>0 && trimmed[i]>='A' && trimmed[i]<='Z'{
			result+=" "
		}
		result+=string(trimmed[i])
	}
	return result
}
func main() {
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
}