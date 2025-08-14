package main

import "fmt"

func QuarterOfAYear(month int) int {
	if month < 1 || month > 12 {
		return -1
	}
	if month <= 3 {
		return 1
	} else if month <= 6 {
		return 2
	} else if month <= 9 {
		return 3
	} else {
		return 4
	}
}

func main() {
	fmt.Println(QuarterOfAYear(2))
	fmt.Println(QuarterOfAYear(5))
	fmt.Println(QuarterOfAYear(9))
	fmt.Println(QuarterOfAYear(11))
	fmt.Println(QuarterOfAYear(13))
	fmt.Println(QuarterOfAYear(-5))
}
