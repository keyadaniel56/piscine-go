package main

import "fmt"

func putstr(s string){
	for _,char:=range s{
		fmt.Printf("%c",char)
	}
}

func main(){
	putstr("hello world")
}