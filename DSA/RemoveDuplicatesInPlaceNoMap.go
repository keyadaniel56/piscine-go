package main

import "fmt"

func RemoveDuplicatesInPlaceNoMap(slice []int) []int {
	j := 0
	for i := 0; i < len(slice); i++ {
		found := false
		for k := 0; k < j; k++ {
			if slice[k] == slice[i] {
				found = true
				break
			}
		}
		if !found {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func main() {
	input := []int{1, 2, 2, 3, 1, 4, 2, 5}
	result := RemoveDuplicatesInPlaceNoMap(input)
	fmt.Println(result) // [1 2 3 4 5]
}
