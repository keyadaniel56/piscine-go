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


func add_prime(n int)int{
	sum:=0
	for i:=2;i<=n;i++{
		if isprime(i){
			sum+=i
		}
	}
	return sum
}


func main(){
	fmt.Println(add_prime(7))
}