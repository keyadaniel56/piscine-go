package main

import "os"

func main(){
	if len(os.Args)!=4{
		os.Stdout.Write([]byte("\n"))
	}

	input:=os.Args[1]
	search:=os.Args[2]
    replace:=os.Args[3]

	if len(search)!=1 || len(replace)!=1{
		os.Stdout.Write([]byte("\n"))
	}

	searchchar:=search[0]
	replacechar:=replace[0]
	result:=[]byte{}

	for i:=0;i<len(input);i++{
		if input[i]==searchchar{
			result=append(result,replacechar)
		}else{
			result=append(result,input[i])
		}
	}
	os.Stdout.Write(result)
	os.Stdout.Write([]byte("\n"))
}