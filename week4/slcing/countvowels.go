package main

import "fmt"

func countvowels(s string)int{
	count:=0
	for _,char:=range s{
		if char=='a' || char=='e' || char=='i' || char=='o' || char=='u' ||char=='A' || char=='E' || char=='I' || char=='O' || char=='U'{
			count++
		}
	}
	return count
}

func main(){
	fmt.Println(countvowels("hello"))
}