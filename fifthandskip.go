package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Remove spaces manually
	s := ""
	for _, c := range str {
		if c != ' ' {
			s += string(c)
		}
	}

	if len(s) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	j := 0
	for _, c := range s {
		result += string(c)
		j++
		if j == 5 {
			result += " "
			j = 0
		}
	}

	result += "\n"
	return result
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}