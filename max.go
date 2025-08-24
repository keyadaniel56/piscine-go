package main

import "fmt"

func Max(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	num := []int{1, 2, 8, 4, 5}
	fmt.Println(Max(num))
}
