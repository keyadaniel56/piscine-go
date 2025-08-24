package main

import "os"

func AtoI(s string) int {
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

func ITOA(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var digit []byte
	for n > 0 {
		d := n % 10
		digit = append(digit, byte(d)+'0')
		n /= 10
	}
	for i, j := 0, len(digit)-1; i < j; i, j = i+1, j-1 {
		digit[i], digit[j] = digit[j], digit[i]
	}
	if isNegative {
		return "-" + string(digit)
	}
	return string(digit)
}
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

func main() {
	if len(os.Args) != 2 {
		os.Stdout.Write([]byte("\n"))
		return
	}
	num := os.Args[1]
	my_num := AtoI(num)
	if my_num <= 0 {
		os.Stdout.Write([]byte("\n"))
		return
	}

	prime := addprimesum(my_num)
	result := ITOA(prime)
	os.Stdout.Write([]byte(result + "\n"))
}
