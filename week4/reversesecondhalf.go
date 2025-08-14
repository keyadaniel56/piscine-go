package main

import "fmt"

func ReverseSecondHalf(str string) string {
	if len(str)==0{
		return "Invalid Output" +"\n"
	}
	second_half:=len(str)/2
	result:=str[second_half:]
	runes:=[]rune(result)
	for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
		runes[i],runes[j]=runes[j],runes[i]
	}
	return string(runes) +"\n"
}

func main() {
	fmt.Print(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(ReverseSecondHalf(""))
	fmt.Print(ReverseSecondHalf("Hello World"))
}