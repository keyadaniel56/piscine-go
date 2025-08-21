package main

import "fmt"

func isCamel(s string)bool{
	if len(s)==0{
		return false
	}
	for i:=0;i<len(s);i++{
		char:=s[i]
		if !(char>='a' && char<='z')&&!(char>='A' && char<='Z'){
			return false
		}
		if i>0 && s[i-1]>='A' && s[i-1]<='Z' && char>='A' && char<='Z'{
			return false
		}
		if i==len(s)-1 && char>='A' && char<='Z'{
			return false
		}
	}
	return true
}


func camelToSnakeCase(s string)string{
	if !isCamel(s){
		return s
	}
	results:=""
	for i:=0;i<len(s);i++{
		char:=s[i]
		if char>='A' && char<='Z'{
			if i>0{

				results+="_"
			}
			results+=string(char)
		}else{
			results+=string(char)
		}
	}
	return results
}

func main(){
	fmt.Println(camelToSnakeCase("helloWorld"))
}