package main

import "fmt"


func Lcm(a,b int)int{
	return (a*b)/gcd(a,b)
}

