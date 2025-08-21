package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	first, second := slice1, slice2
	results := []int{}
	if len(first) > len(second) {
		for i := len(first) - 1; i >= len(second); i-- {
			results = append(results, first[i])
		}
	} else if len(second) > len(first) {
		for i := len(second) - 1; i >= len(first); i-- {
			results = append(results, second[i])
		}
	}
	m := len(first)
	if len(second) < m {
		m = len(second)
	}
	for i := m - 1; i >= 0; i-- {
		results = append(results, first[i])
		results = append(results, second[i])
	}
	return results
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
