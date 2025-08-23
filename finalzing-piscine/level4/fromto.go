package main

import "fmt"
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var digits []byte
	for n > 0 {
		d := n % 10
		digits = append(digits, byte(d)+'0')
		n /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	if isNegative {
		return "-" + string(digits)
	}
	return string(digits)
}
func FromTo(from int, to int) string {
	if from<0 || from> 99 || to<0 || to>99{
		return "Invalid\n"
	}
	result:=""
	step:=1

	if from>to{
		step=-1
	}
	for i:=from;;i+=step{
		if i<10{
			result+="0"
		}
		result+=itoa(i)
		if i==to{
			break
		}
		result+=","
	}
	return result+"\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}