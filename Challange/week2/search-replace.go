package main

import "os"

func search_replace(input,search,replace string)string{
	result:=""
	if len(search)!=1 || len(replace)!=1{
		return input
	}
	s:=search[0]
	r:=replace[0]
	for i:=0;i<len(input);i++{
		if input[i]==s{
			result+=string(r)
		}else{
			result+=string(input[i])
		}
	}
	return result
}
func main(){
	fmt.Println(search_replace("hello","h","o"))
}