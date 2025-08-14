package main

import "fmt"

func PrintRange(start, end int) {
	if start < end {
		for i:=start;i<end;i++{
			fmt.Print(i," ")
		}
		fmt.Println(end)
	}else if start > end {
		for i:=start-1;i>end;i--{
			fmt.Print(i," ")
		}
		fmt.Println(end)
	}else{
		if start!=10{
			fmt.Println(start)
		}else{
			fmt.Println()
		}
	}
}

func main() {
	PrintRange(1, 10)
	PrintRange(10, 1)
	PrintRange(1, 1)
	PrintRange(10, 10)
	PrintRange(0, 9)
	PrintRange(-1, -10)
}
