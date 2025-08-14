package main

import "fmt"

func RemoveDuplicate(slice []int) []int {
	seen:=make(map[int]bool)
	var result [] int
	for _,v:=range slice{
		if !seen[v]{
			seen[v]=true
			result=append(result, v)
		}
	}
	return result
}


func main() {
	fmt.Println(RemoveDuplicate([]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(RemoveDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(RemoveDuplicate([]int{}))
}