package main

import "fmt"

func RotateSlice(slice []int, k int) []int {
	if len(slice) == 0 {
		return slice
	}
	k = k % len(slice)
	return append(slice[:k], slice[k:]...)
}

func main() {
	fmt.Println(RotateSlice([]int{1, 2, 3, 4, 5}, 2)) // [3 4 5 1 2]

}
