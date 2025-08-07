package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	result := 0
	start := 0
	sign := 1

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}
	return result * sign
}

func mulTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as argument")
		return
	}
	n := Atoi(os.Args[1])
	mulTable(n)
}
