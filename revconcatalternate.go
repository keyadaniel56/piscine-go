package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1:=len(slice1)
	len2:=len(slice2)
	maxlen:=len1
	result:=make([]int,0,len1+len2)

	if len2>maxlen{
		maxlen=len2
	}
	for i:=maxlen;i>=0;i--{
		if i<len1{
			result=append(result, slice1[i])
		}
		if i<len2{
			result=append(result, slice2[i])
		}
	}
	return result
}