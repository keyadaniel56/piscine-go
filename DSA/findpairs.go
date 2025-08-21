package main

import (
	"fmt"
	"os"
)

// --- Custom Atoi (returns (int, bool)) ---
func Atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	start := 0
	result := 0
	sign := 1

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}
	return result * sign, true
}

// --- Trim spaces ---
func removetrail(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && s[start] == ' ' {
		start++
	}
	for end >= start && s[end] == ' ' {
		end--
	}
	if start > end {
		return ""
	}
	return s[start : end+1]
}

// --- Split by commas ---
func separatebycomma(s string) []string {
	var parts []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			parts = append(parts, removetrail(word))
			word = ""
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		parts = append(parts, removetrail(word))
	}
	return parts
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	// --- Validate array format ---
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}
	arrStr = arrStr[1 : len(arrStr)-1] // remove [ and ]

	parts := separatebycomma(arrStr)
	nums := []int{}
	for _, p := range parts {
		if p == "" {
			fmt.Println("Invalid input.")
			return
		}
		num, ok := Atoi(p)
		if !ok {
			fmt.Printf("Invalid number: %s\n", p)
			return
		}
		nums = append(nums, num)
	}

	// --- Parse target ---
	target, ok := Atoi(removetrail(targetStr))
	if !ok {
		fmt.Println("Invalid target sum.")
		return
	}

	// --- Find pairs ---
	pairs := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	// --- Output ---
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}
