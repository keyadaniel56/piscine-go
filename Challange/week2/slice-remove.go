package main

import "fmt"


func SliceRemove(slice []int , num int) []int {
	for i,v:=range slice{
		if v==num{
			return append(slice[:i],slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	fmt.Println(SliceRemove([]int{1, 2, 3}, 2))
	fmt.Println(SliceRemove([]int{4, 3}, 4))
	fmt.Println(SliceRemove([]int{}, 1))

}