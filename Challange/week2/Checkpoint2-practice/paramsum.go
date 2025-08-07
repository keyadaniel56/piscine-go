package main

import ("fmt"
   "os")

func Atoi(s string)int{
	start:=0
	result:=0
	sign:=1

	if s[0]=='-'{
		sign=-1
		start=1
	}
	if s[0]=='+'{
		start=1
	}

	for i:=start;i<len(s);i++{
		if s[i]<'0' || s[i]>'9'{
			return 0
		}
		digit:=int(s[i]-'0')
		result=result*10+digit
	}
	return result*sign
}



func param_sum()int{
	if len(os.Args)<2{
		return 0
	}
	n:=Atoi(os.Args[1])
	sum:=0
	for i:=0;i<=n;i++{
		sum+=i
	}
	return sum
}



func main(){
	fmt.Println(param_sum())
}