package main

import "fmt"

func FtRange(start, end int) []int {
	size := 0
	if start <= end {
		size = end - start + 1
	} else {
		size = start - end + 1
	}
	result := make([]int, size)
	if start <= end {
		for i := 0; i < size; i++ {
			result[i] = start + i
		}
	} else {
		for i := 0; i < size; i++ {
			result[i] = start - i
		}
	}
	return result
}
