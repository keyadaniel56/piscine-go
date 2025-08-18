package main

import "fmt"

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n==0{
		return 1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		n = n / base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))  // 3
	fmt.Println(DigitLen(100, 2))   // 7
	fmt.Println(DigitLen(-100, 16)) // 2
	fmt.Println(DigitLen(100, -1))  // -1
}
