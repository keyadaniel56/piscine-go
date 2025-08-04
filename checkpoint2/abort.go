
// Instructions
// Write a function that returns the median of five int arguments.

// Expected function
// func Abort(a, b, c, d, e int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	middle := piscine.Abort(2, 3, 8, 5, 7)
// 	fmt.Println(middle)
// }
// And its output :

// $ go run .
// 5
// $





package main

import "fmt"

func sortintegertable(table [] int){
	for i:=0;i<len(table)-1;i++{
		for j:=0;j<len(table)-i-1;j++{
			if table[j]>table[j+1]{
				table[j],table[j+1]=table[j+1],table[j]
			}
		}
	}
}



func Abort(a, b, c, d, e int) int {
args:=[]int{a,b,c,d,e}
sortintegertable(args)
return args[2]
}


func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}



