package main

import "fmt"

func printhex(n int){
	if n==0{
		return "0"
	}

	hexchar:="0123456789abcdef"
	result:=""
	for n>0{
		reminder:=n%16
		result=string(hexchar[reminder]+result)
		reminder/=16
	}
	fmt.Println(result)
}