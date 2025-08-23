package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NextPime(n int) int {
	candidate := n + 1
	for {
		if isPrime(candidate) {
			return candidate
		}
		candidate++
	}
}
func main() {
	fmt.Println(NextPime(7))
}
