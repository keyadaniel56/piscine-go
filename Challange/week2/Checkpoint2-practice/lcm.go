package main

import "fmt"

func gcd(a,b int)int{
	for b!=0{
		a,b=b,b%a
	}
	return a 
}


func lcm(a,b int)int{
	return (a*b)/gcd(a,b )
}


func main(){
	fmt.Println(lcm(12,18))
}