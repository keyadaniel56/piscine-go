package main

import "fmt"

func isprime(n int)bool{
	for i:=2;i*i<=n;i++{
		if n%i==0{
			return false
		}
	}
	return true
}


func prevprime(n int)int{
	if  n<2{
		return 0
	}
	for n>=2{
		if isprime(n){
			return n
		}
		n+=1
	}
	return 0
}


func main(){
	fmt.Println(prevprime(12))
}