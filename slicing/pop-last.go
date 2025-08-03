package main

import "fmt"

func main(){
	s := []int{5, 10, 15, 20, 25, 30}
	fmt.Println(s[:len(s)-1])
}