package main

import "fmt"

func ConcatAlternate(a, b []int) (r []int) {
	if len(b)>len(a){
		a,b=b,a
	}
	for i:=range a{
		r = append(r, a[i])
		if i<len(b){
			r = append(r, b[i])
		}
	}
	return
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
