package main

import "fmt"


func main(){
	for i:='z';i>='a';i--{
		if i%2==0{
            fmt.Printf("%c",i)
		}else{
			fmt.Printf("%c",i-32)
		}
	}
}