package main

import "fmt"

func print_numbers(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
	}
}

func main() {
	print_numbers(9)
}
