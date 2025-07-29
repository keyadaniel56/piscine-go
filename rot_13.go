package main

import "os"

func rot_13(s string)string{
	b:=[]rune(s)
	for i,char:=range b{
		if char>='a' && char<='z'{
			b[i]=(char-'a'+13)%26+'a'
		}else if char>='A' && char<='Z'{
			b[i]=(char-'A'+13)%26+'A'
		}
	}
	return string(b)
}

func main(){
	if len(os.Args)!=1{
		os.Stdout.Write([]byte("\n"))
	}
	result:=os.Args[1]
	encoded:=rot_13(result)
	os.Stdout.Write([]byte(encoded+"\n"))
}