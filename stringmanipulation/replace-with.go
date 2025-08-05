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
		}else if  word!=""{
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
	trimmed:=TrimFields(s)
	words:=split(trimmed)
	result:=""
	for i:=0;i<len(words);i++{
		result+=string(words[i])
		if i!=len(words)-1{
			result+=separator
		}
	}
	return result
}


func capitalize(s string)string{
	runes:=[]rune(s)
	if len(s)==0{
		return s
	}
   for i:=0;i<len(runes);i++{
	if (i==0 || runes[i-1]==' ') && (runes[i]>='a' && runes[i]<='z'){
		runes[i]-=32
	}
   }
   return string(runes)
}


    func tolower(s string) string {
	runes := []rune(s)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = char + ('a' - 'A')
		}
	}
	return string(runes)
}


func snakecase(s string)string{
words:=tolower(s)
result:=separator(words,"_")
return result
}


func camelCase(s string) string {
	words := split(s)        // split input into words
	for i := range words {
		words[i] = tolower(words[i]) // lowercase each word
		if i != 0 {
			words[i] = capitalize(words[i]) // capitalize first letter except first word
		}
	}
	result := ""
	for _, w := range words {
		result += w
	}
	return result
}



func main(){
	word:="hello World i love coding"
	fmt.Println(tolower(word))
	fmt.Println(capitalize(word))
	fmt.Println(split(word))
	fmt.Println(separator(word , "_"))
	fmt.Println(snakecase("hello world i love coding"))
	fmt.Println(camelCase(word))
}