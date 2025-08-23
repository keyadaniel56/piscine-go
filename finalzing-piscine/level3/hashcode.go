package main

import "fmt"

func hashcode(dec string) string {
	size := len(dec)
	result := ""
	for i := 0; i < size; i++ {
		val := int(dec[i]+byte(size)) % 127
		if val < 33 {
			val += 33
		} else {
			result += string(val)
		}
	}
	return result
}

func main() {
	fmt.Println(hashcode("A"))
	fmt.Println(hashcode("AB"))
	fmt.Println(hashcode("BAC"))
	fmt.Println(hashcode("Hello World"))
}
