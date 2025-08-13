package main

import "fmt"
func VowelsIndex(str string) []int {
	var result [] int
	vowels:=map[rune]bool{'a':true,'e':true,'i':true,'o':true,'u':true,
                          'A':true,'E':true,'I':true,'O':true,'U':true}
for i,char:=range str{
	if vowels[char]{
		result=append(result, i)
	}
}
return result
}

func main() {
	res := []string{"student", "hello Iyan","bcdfgh", "wOrld", "","AAEO$o;jw"}
	for _, i := range res {
		fmt.Println(VowelsIndex(i))
	}
}