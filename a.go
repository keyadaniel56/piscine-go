package main

import "fmt"

func FindA(s string) {
	if len(s) == 0 {
		fmt.Println("a")
	}
	for _, char := range s {
		if char == 'a' {
			fmt.Println("a")
			return
		} else {
			fmt.Println()

		}
	}
}

func main() {
	FindA("vvzv")

}
