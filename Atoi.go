package main

import "fmt"

func Atoi(s string) int {
	start := 0
	sign := 1
	result := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}

func main() {
	fmt.Println(Atoi("-123n"))
}
