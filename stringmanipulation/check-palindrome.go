package main

import "fmt"

// isPalindrome checks if s is a palindrome (case-sensitive, no external packages).
func isPalindrome(s string) bool {
    start, end := 0, len(s)-1

    for start < end {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }
    return true
}

func main() {
    test := "madam"
    fmt.Printf("%q palindrome? %v\n", test, isPalindrome(test))

    test2 := "racecar"
    fmt.Printf("%q palindrome? %v\n", test2, isPalindrome(test2))
}


func palindrome(s string)bool{
   start,end:=0,len(s)-1

   for start<end {
    if s[start]!=s[end]{
        return false
    }
    start ++
    end--
   }
   return true
}