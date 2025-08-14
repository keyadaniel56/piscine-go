package main

import "fmt"

func SliceAdd(slice []int , num int) []int {
	return append(slice,num)
}

func main() {
	fmt.Println(SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(SliceAdd([]int{}, 4))
}