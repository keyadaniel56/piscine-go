package main

import "fmt"

func swap(a,b *int){
	*a,*b=*b,*a
}

func main(){
	a:=10
	b:=12
	swap(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}