package main

import "fmt"

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func main() {
	// Test cases
	fmt.Println(Sqrt(16)) // Should return 4
	fmt.Println(Sqrt(25)) // Should return 5
	fmt.Println(Sqrt(10)) // Should return 0 (not a perfect square)
	fmt.Println(Sqrt(0))  // Should return 0 (square root of 0 is 0)
	fmt.Println(Sqrt(-4)) // Should return 0 (no real square root for negative numbers)
}
