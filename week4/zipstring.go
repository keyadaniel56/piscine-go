package main

import "fmt"

func countDuplicate(s string, i int)int{
	if i<0 || i>=len(s){
		return 0
	}

	charpos:=rune(s[i])
	count:=0
	for _,char:=range s{
		if char==charpos{
			count++
		}
	}
	return count
}


func itoa(n int)string{
	if n==0{
		return "0"
	}
	isNegative:=false
	if n<0{
		isNegative=true
		n=-n
	}

	var digits [] rune
	for n>0{
	d:=rune(n%10+'0')
	digits=append(digits, d)
	n/=10
}
	for i,j:=0,len(digits)-1;i<j;i,j=i+1,j-1{
		digits[i],digits[j]=digits[j],digits[i]
	}
	if isNegative{
		return "-" + string(digits)
	}
	return string(digits)
}


func ZipString(s string) string {
	var result string
	i:=0
	for i<len(s){
		counter:=countDuplicate(s,i)
		result=result+itoa(counter)+string(s[i])
		i+=int(counter)
	}
	return result
}
func main(){
	fmt.Println(itoa(123))
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}