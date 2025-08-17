package main

import "fmt"

func SaveAndMiss(arg string,num int)string{
	if num<=0 || len(arg)==0{
		return arg
	}

	if len(arg)<num{
		return arg
	}
	if len(arg)<2*num{
		return arg[:num]
	}
	return arg[:num]+SaveAndMiss(arg[2*num:],num)
}