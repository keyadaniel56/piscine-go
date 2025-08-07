package main 

import "fmt"


func AtoiBase(s string, base int )int{
	start:=0
	sign:=1
	result:=0
    if base<2 || base > 16{
		return 0
	}

	if len(s)==0{
		return 0
	}
	if s[0]=='-'{
		sign=-1
		start=1
	}
	if s[0]=='+'{
		start=1
	}
    
	for i:=start;i<len(s);i++{
	ch:=s[i]
	var digit int 

	if ch>='0' && ch<='9'{
		digit=int(ch-'0')
	}else if ch>='a' && ch<='f'{
		digit=int(ch-'a'+10)
	}else if ch>='A' && ch<='F'{
		digit=int(ch-'A'+10)
	}else{
		break
	}
	if digit>=base{
		break
	}
	result=result*base + digit
	}
	return result*sign
}


func main(){
	fmt.Println(AtoiBase("-123b",16))
}