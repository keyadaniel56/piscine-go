package main

import "fmt"

func gcd(a, b int) int {
	if a == 0 {
		return a
	}
	if b == 0 {
		return b
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(gcd(12, 4))
}
