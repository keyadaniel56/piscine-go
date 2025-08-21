package main

import "fmt"

func RemoveDuplicatesInPlace(slice []int) []int {
	seen := make(map[int]bool)
	j := 0
	for i := 0; i < len(slice); i++ {
		if !seen[slice[i]] {
			seen[slice[i]] = true
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func main() {
	input := []int{1, 2, 2, 3, 1, 4, 2, 5}
	result := RemoveDuplicatesInPlace(input)
	fmt.Println(result) // [1 2 3 4 5]
}
