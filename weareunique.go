package main

import "fmt"

func WeAreUnique(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return -1
	}

	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)

	for _, c := range s1 {
		if c != ' ' {
			set1[c] = true
		}
	}
	for _, c := range s2 {
		if c != ' ' {
			set2[c] = true
		}
	}

	count := 0

	for c := range set1 {
		if !set2[c] {
			count++
		}
	}
	for c := range set2 {
		if !set1[c] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", " "))   // 2, ignoring the space
	fmt.Println(WeAreUnique("foo", "boo")) // 2
	fmt.Println(WeAreUnique("", ""))       // -1
	fmt.Println(WeAreUnique("abc", "def")) // 6
}
