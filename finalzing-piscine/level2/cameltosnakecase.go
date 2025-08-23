package main

import "fmt"

func isCamel(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] < 'a' || s[0] > 'z' {
		return false
	}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
		if i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z' && char >= 'A' && char <= 'Z' {
			return false
		}
		if i == len(s)-1 && char >= 'A' && char <= 'Z' {
			return false
		}
	}
	return true
}

func camel_to_snake(s string)string{
	if !isCamel(s){
		return s
	}
	result:=[]rune{}
	for i,r:=range s{
		if (i>0 && r>='A' && r<='Z'){
			result=append(result, '_')
			result = append(result, r)
		}else{
			result = append(result, r)
		}
	}
 return string(result)
}
func main() {
	fmt.Println( camel_to_snake("helloWorld"))      // true
	fmt.Println( camel_to_snake("hello world"))     // false (space)
	fmt.Println( camel_to_snake("HelloWorld"))      // false (starts with uppercase)
	fmt.Println( camel_to_snake("helloWorldZ"))     // false (ends with uppercase)
	fmt.Println( camel_to_snake("helloWWWorld"))    // false (consecutive uppercase)
	fmt.Println( camel_to_snake("hello123"))        // false (contains digits)
	fmt.Println( camel_to_snake(""))                // false (empty)
	fmt.Println( camel_to_snake("camelCaseString")) // true
}
