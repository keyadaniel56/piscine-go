package main

import "fmt"

func Unmatch(a []int) int {
	seen:=make(map[int]int)
	for _,v:=range a{
		seen[v]++
	}
	for _,v:=range a{
		if seen[v]==1{
			return v
		}
	}
	return -1
}


func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch) // Expected output: 4
}