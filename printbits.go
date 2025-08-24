package main

import "fmt"

func printBits(n uint) {
	start := false
	for i := 31; i >= 0; i-- {
		if n&(1<<i) != 0 {
			start = true
		} else if start {
			fmt.Println("0")
		}
	}
	if !start {
		fmt.Print("0")
	}
	fmt.Println()
}
