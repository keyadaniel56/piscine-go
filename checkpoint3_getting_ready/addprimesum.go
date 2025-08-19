package main

import (
	"os"
)

func isprime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func addprimesum(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isprime(i) {
			sum += i
		}
	}
	return sum
}

func Atoi(s string) int {
	start := 0
	result := 0
	sign := 1
	if s[0] == '-' {
		sign = -1
		start = 1
	}
	if s[0] == '+' {
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

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var digits []byte
	for n > 0 {
		d := n % 10
		digits = append(digits, byte(d)+'0')
		n /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	if isNegative {
		return "-" + string(digits)
	}
	return string(digits)
}

func main() {
	if len(os.Args) < 2 {
		os.Stdout.Write([]byte("\n"))
	}
	input := os.Args[1]
	num := Atoi(input)
	result := addprimesum(num)
	output := Itoa(result)
	os.Stdout.Write([]byte(output))
	os.Stdout.Write([]byte("\n"))
}
