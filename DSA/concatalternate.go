package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	results := []int{}
	first, second := slice1, slice2
	if len(slice2) > len(slice1) {
		first, second = slice2, slice1
	}
	for i := 0; i < len(second); i++ {
		results = append(results, first[i])
		results = append(results, second[i])
	}
	if len(first) > len(second) {
		results = append(results, first[len(second):]...)
	}
	return results
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
