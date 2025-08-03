package main

func isprime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func addprime(n int) int {
	sum := 0
	for i := 2; i<= n; i++ {
		if isprime(i) {
			sum+=i
		}
	}
	return sum
}

func main() {
	println(isprime(7))  // true
	println(isprime(8))  // false
	println(addprime(7)) // 4 (primes are 2, 3, 5, 7)
}
