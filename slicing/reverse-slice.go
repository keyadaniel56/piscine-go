package main

import (
	"fmt"
)

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before reverse:", arr)

	reverse(arr)
	fmt.Println("After reverse: ", arr)

	// Test with even number of elements
	arr2 := []int{10, 20, 30, 40}
	fmt.Println("Before reverse:", arr2)

	reverse(arr2)
	fmt.Println("After reverse: ", arr2)
}
