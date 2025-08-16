package main

import "fmt"

func FishAndChips(n int) string {
	if n%2==0 && n%3==0{
		return "fish and chips"
	}else if n%2==0{
		return "chips"
	}else if n%3==0{
		return "fish"
	}
		return "error: non divisible"
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
	fmt.Println(FishAndChips(7))
}