package main

import "fmt"

func iscamelCase(s string)bool{
	if len(s)==0{
		return false
	}
	for i:=0;i<len(s);i++{
		char:=s[i]
		//must not contain digits
		if !(char>='a' && char<='z') && !(char>='A' && char<='Z'){
			return false
		}
		//no tweo consecutive uppercase
		if i>0 && s[i-1]>='A' && s[i-1]<='Z' && char>='A' && char<='Z'{
			return false
		}
		//must not end with uppercase
		if i==len(s)-1 && char>='A' && char<='Z'{
			return false
		}
	}
	return true
}


func camelToSnakeCase(s string)string{
	if !iscamelCase(s){
		return s
	}
	result:=""
	for i:=0;i<len(s);i++{
		char:=s[i]
		if char>='A' && char<='Z'{
			if i>0{
				result+="_"
			}
			result+=string(char)
		}else{
			result+=string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(camelToSnakeCase("HelloWorld"))
	fmt.Println(camelToSnakeCase("helloWorld"))
	fmt.Println(camelToSnakeCase("camelCase"))
	fmt.Println(camelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(camelToSnakeCase("camelToSnakeCase"))
	fmt.Println(camelToSnakeCase("hey2"))
}