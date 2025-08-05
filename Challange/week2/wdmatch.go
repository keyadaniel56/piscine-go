package main

import "fmt"


func Wdmatch(s , s2 string)string{
	i,j:=0,0
	for i<len(s) && j<len(s2){
		if s[i]==s2[j]{
			i++
		}
		j++
	}
	if i==len(s){
		return s
	}
	return ""
}

func main() {
	fmt.Println( Wdmatch("faya", "fgvvfdxcacpolhyghbreda")) // should print "faya"
	fmt.Println( Wdmatch("foo", "boo"))                      // should print ""
	fmt.Println( Wdmatch("abc", "def"))                      // should print ""
}