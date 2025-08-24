package main

import "fmt"


func SaveAndMiss(arg string, num int) string {
	if num<=0{
		return arg
	}
	result:=""
	save:=true
	for i:=0;i<len(arg);{
		if save{
			end:=i+num
			if end>len(arg){
				end=len(arg)
			}
			result+=arg[i:end]
		}
		i+=num
		save=!save
	}
	return result
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}