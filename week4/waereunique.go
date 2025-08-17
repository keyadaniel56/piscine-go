package main

import "fmt"

func WeAreUnique(str1 , str2 string) int {
	if len(str1)==0 && len(str2)==0{
		return -1
	}
	set1:=make(map[rune]bool)
	set2:=make(map[rune]bool)

	for _,char:=range str1{
		set1[char]=true
	}
	for _,char:=range str2{
		set2[char]=true
	}

	seen:=make(map[rune]bool)
	count:=0
	for char:=range set1{
		if !set2[char] && !seen[char]{
			count++
			seen[char]=true
		}
	}
	for char:=range set2{
		if !set1[char] && !seen[char]{
			count++
			seen[char]=true
		}
	}
	return count
}


func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}