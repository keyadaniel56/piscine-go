package main

import "fmt"

func SwapLast(slice []int) []int {
	n:=len(slice)
	if n<2{
		return slice
	}
	slice[n-1],slice[n-2]=slice[n-2],slice[n-1]
	return slice
}

func main() {
	fmt.Println(SwapLast([]int{1, 2, 3, 4}))
	fmt.Println(SwapLast([]int{3, 4, 5}))
	fmt.Println(SwapLast([]int{1}))
}