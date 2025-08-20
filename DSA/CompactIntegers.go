package main

import "fmt"

func CompactIntegers(slice []int) []int {
	var results []int
	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			results = append(results, slice[i])
		}
	}
	return results
}

func main() {
	input := []int{0, 1, 0, 2, 3, 0, 4, 5}
	result := CompactIntegers(input)
	fmt.Println(result)

}
