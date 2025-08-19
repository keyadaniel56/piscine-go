package main

import "fmt"

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return true
	}
	pos := 0
	for pos < len(arr)-1 {
		step := int(arr[pos])
		if step == 0 {
			return false
		}
		pos += step
		if pos == len(arr)-1 {
			return true
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
