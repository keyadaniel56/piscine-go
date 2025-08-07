package main

import "fmt"

func hidenp(s1,s2 string)string{
	runes:=[]rune(s1)
	i:=0
	for _,char:=range s2{
		if i<len(s2) && char==runes[i]{
			i++
		}
		if i==len(runes){
			return "1"
		}
	}
	return "0"
}


func main(){
	fmt.Println(hidenp("fgex.;","tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"))
}


