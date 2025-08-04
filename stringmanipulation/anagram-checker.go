
package main


 import "fmt"

func Anagram(s1,s2 string)bool{
	var is_seen [] bool

	if len(s1)!=len(s2){
		return false
	}
	for _,char1:=range s1{
		for _,char2:=range s2{
			if char1==char2{
				is_seen=append(is_seen,true)
			}
		}
	}

	if len(s1)==len(is_seen){
		return true
	}else{
		return false
	}
}


func main(){
	result:=Anagram("silent","listen")
	fmt.Println(result)
}