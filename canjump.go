package main

import "fmt"

func Canjump(arr []uint)bool{
	if len(arr)==0{
		return false
	}
	if len(arr)==1{
		return true
	}
	pos:=0
	for pos<len(arr)-1{
		step:=int(arr[pos])
		if step==0{
			return false
		}
		pos+=step
		if pos==len(arr)-1{
			return true
		}
	}
	return false
}