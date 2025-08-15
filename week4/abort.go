package main

import (
	"fmt"
	
)

func sort(arr[]int){
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
}


 func Abort(a, b, c, d, e int) int {
	arr:=[]int{a,b,c,d,e}
	sort(arr)
	return arr[2]

}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}