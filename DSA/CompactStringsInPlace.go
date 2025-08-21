package main

import "fmt"

func CompactStringsInPlace(slice []string) []string {
	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func main() {
	input := []string{"hello", "", "world", "", "", "go"}
	result := CompactStringsInPlace(input)
	fmt.Println(result)

}
