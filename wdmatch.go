package main

import "fmt"

func wdmatch(s1, s2 string) bool {
	i := 0
	for _, c := range s2 {
		if i < len(s1) && byte(c) == s1[i] {
			i++
		}
	}
	return i == len(s1)
}
