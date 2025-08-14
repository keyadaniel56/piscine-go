package main

import "fmt"

func SwapFirst(slice []int) []int {
	n:=len(slice)
	if n<2{
		return slice
	}
	slice[0],slice[1]=slice[1],slice[0]
	return slice
}

func main() {
	fmt.Println(SwapFirst([]int{1, 2, 3, 4}))
	fmt.Println(SwapFirst([]int{3, 4, 6}))
	fmt.Println(SwapFirst([]int{1}))
}
