package main

import "fmt"

func Split(s, sep string) []string {
	var word[]string
	if sep==""{
		return []string{s}
	}
	start:=0
	for i:=0;i<len(s)-len(sep);{
		if i+len(sep)<=len(s) && s[i:i+len(sep)]==sep{
			if start<i{
				word=append(word,s[start:i])
			}
			start=i+len(sep)
			i=start
		}else{
			i++
		}
	}
	if start<len(s){
		word=append(word, s[start:])
	}
	return word
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n",Split(s, "HA"))
}