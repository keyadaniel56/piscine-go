package main

import "os"

// Convert string to int (basic version, assumes valid positive number)
func atoi(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}

// Convert int to string
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string('0'+n%10) + result
		n /= 10
	}
	return result
}

// Print string to stdout
func print(s string) {
	os.Stdout.Write([]byte(s))
}

func main() {
	if len(os.Args) != 2 {
		print("\n")
		return
	}

	n := atoi(os.Args[1])

	for i := 1; i <= 9; i++ {
		print(itoa(i) + " x " + itoa(n) + " = " + itoa(i*n) + "\n")
	}
}
