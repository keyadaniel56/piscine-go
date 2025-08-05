package main

import "fmt"


func ZipString(s string) string {
	count:=1
	result:=""

	for i:=1;i<len(s);i++{
		if s[i]==s[i-1]{
			count++
		}else{
			result+=string(s[i-1])
			if count>1{
				result+=fmt.Sprintf("%d",count)
			}
			count=1
		}
	}
	result+=string(s[len(s)-1])
	if count>1{
		result+=fmt.Sprintf("%d",count)
	}
	return result
}


func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}