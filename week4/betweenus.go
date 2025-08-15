package main

import "fmt"

func BetweenUs(num, min, max int) bool {
	// Your code here
	if min > max {
		return false
	}
	return num >= min && num <= max

}

func main() {
	fmt.Println(BetweenUs(1, 2, 3))
	fmt.Println(BetweenUs(1, 1, 3))
	fmt.Println(BetweenUs(1, 3, 3))
	fmt.Println(BetweenUs(1, 1, 1))
	fmt.Println(BetweenUs(1, 2, 1))
	fmt.Println(BetweenUs(-1, -10, 0))
}
