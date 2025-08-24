package main

import "fmt"

func ispowerof2(n int) int {
	if n == 0 {
		return 0
	}
	if n&(n-1) == 0 {
		return 1
	}
	return 0
}
