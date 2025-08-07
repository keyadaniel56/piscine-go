package main

import "fmt"


func ft_range(start,end int)[]int{
	length:=0
	if end>=start{
		length=end-start+1
	}else{
		length=start-end+1
	}
	arr:=make([]int,length)

	if start<=end{
		for i:=0;i<length;i++{
			arr[i]=start+i
		}
	}else{
		for i:=0;i<length;i++{
			arr[i]=start-i
		}
	}
	return arr
}