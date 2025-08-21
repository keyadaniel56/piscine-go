package main

import "fmt"

func CompactIntegersInPlace(slice []int) []int {
	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func main() {
	input := []int{0, 1, 0, 2, 3, 0, 4, 5}
	result := CompactIntegersInPlace(input)
	fmt.Println(result) // [1 2 3 4 5]
}
