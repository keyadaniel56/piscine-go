package main

import "fmt"

func FindZ(s string){
	if len(s)==0{
		fmt.Println("z")
	}
	for _,char:=range s{
		if char=='z'{
			fmt.Println("z")
			return
		}
	}
	fmt.Println("\n")
}

func main(){
	FindZ("")
}