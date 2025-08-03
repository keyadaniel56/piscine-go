package main

import "fmt"

func max(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	numbers := []int{3, 6, 2, 9, 4}
	fmt.Println(max(numbers)) // Output: 9
}
