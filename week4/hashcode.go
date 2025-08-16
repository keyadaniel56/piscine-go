package main

import "fmt"

func HashCode(dec string) string {
	size := len(dec)
	result := ""
	for i := 0; i < size; i++ {
		val := (int(dec[i]) + size) % 127
		if val < 33 {
			val += 33
		}
		result += string(val)
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))           // B
	fmt.Println(HashCode("AB"))          // CD
	fmt.Println(HashCode("BAC"))         // EDF
	fmt.Println(HashCode("Hello World")) // Spwwz+bz}wo
}
