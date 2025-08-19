package main

import "fmt"

func SaveandMiss(arg string,num int)string{
	if len(arg)==0 || num<=0{
		return arg
	}
	result:=""
	save:=true
	for i:=0;i<len(arg);i+=num{
		end:=i+num
		if end>len(arg){
			end=len(arg)
		}
		if save{
			result+=arg[i:end]
		}
		save=!save
	}
	return result
}

func main() {
	fmt.Println(SaveandMiss("123456789", 3))
	fmt.Println(SaveandMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveandMiss("", 3))
	fmt.Println(SaveandMiss("hello you all ! ", 0))
	fmt.Println(SaveandMiss("what is your name?", 0))
	fmt.Println(SaveandMiss("go Exercise Save and Miss", -5))
}