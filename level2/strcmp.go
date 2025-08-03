package main

import "fmt"

func strcomp(s1,s2 string)int{
	if len(s1)>len(s2){
		return 1
	}else if len(s1)<len(s2){
		return -1
	}
	return 0
}

func main(){
	fmt.Println(strcomp("abc","abc"))
}