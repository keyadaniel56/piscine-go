package main

import "fmt"

func my_atoi(s string)int{
	result:=0
	sign:=1
	start:=0

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

func my_itoa(n int)string{
	if n==0{
		return "0"
	}
	isNegative:=false
	if n<0{
		isNegative=true
		n=-n
	}
	var digits [] byte
	for n>0{
		d:=n%10
		digits=append(digits, byte(d)+'0')
		n/=10
	}
	for i,j:=0,len(digits)-1;i<j;i,j=i+1,j-1{
		digits[i],digits[j]=digits[j],digits[i]
	}
	if isNegative{
		return "-"+string(digits)
	}
	return string(digits)
}


func NotDecimal(s string)string{
	if len(s)==0{
		return "\n"
	}
	dot:=-1
	for i,c:=range s{
		if c=='.'{
			dot=i
			break
		}
	}
	if dot==-1{
		return s+"\n"
	}
	newstr:=s[:dot]+s[dot+1:]
	for i,c:=range newstr{
		if i==0 && c=='-'{
			continue
		}
		if c<'0' || c>'9'{
			return s+"\n"
		}
	}
	n:=my_atoi(newstr)
	return my_itoa(n)+"\n"
}


func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}