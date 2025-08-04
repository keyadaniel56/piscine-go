package main

import "fmt"
func CountActiveBits(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 { // check if the last bit is 1
			count++
		}
		n = n >> 1 // shift bits to the right
	}
	return count
}


func main(){
	fmt.Println(CountActiveBits(3))
}



// Instructions
// Write a function, ActiveBits, that returns the number of active bits (bits with the value 1) in the binary representation of an integer number.

// Expected function
// func ActiveBits(n int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.ActiveBits(7))
// }
// And its output :

// $ go run .
// 3
// $