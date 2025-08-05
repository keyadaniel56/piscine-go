package main

import "fmt"

func WeAreUnique(str1 , str2 string) int {
   set1:=make(map[rune]bool)
   set2:=make(map[rune]bool)

   if str1=="" && str2==""{
      return -1
   }

   for _,char:=range str1{
      set1[char]=true
   }
   for _,char:=range str2{
      set2[char]=true
   }

   count:=0
   for _,char:=range str1{
      if !set2[char]{
         count++
      }
   }
   for _,char:=range str2{
      if !set1[char]{
         count++
      }
   }
   return count
}





func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}