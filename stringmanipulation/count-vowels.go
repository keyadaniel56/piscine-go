package main

import "fmt"


func count_vowels(s string)int{

	count:=0
	for _,char:=range s{
		if char=='a' || char=='e' || char=='i' || char=='o' || char=='u'{
			count++
		}
		
	}
	return count
}
func main(){
fmt.Println(count_vowels("aeiou"))
}