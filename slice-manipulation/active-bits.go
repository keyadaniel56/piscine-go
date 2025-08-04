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