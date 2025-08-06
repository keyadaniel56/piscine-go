package main

import "fmt"


func SecondHalf(slice []int) []int {
	start:=len(slice)/2
	return slice[start:]
}


func main() {
	fmt.Println(SecondHalf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(SecondHalf([]int{1, 2, 3}))
}