package main

import "fmt"

func isPrime(n int)bool{
	if n<2{
		return false
	}
	for i:=2;i*i<=n;i++{
		if n%i==0{
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for i:=nb;i>=0;i--{
		if isPrime(i){
			return i
		}
	}
	return 0
}
func main() {
	fmt.Println(FindPrevPrime(500))
	fmt.Println(FindPrevPrime(400))
}