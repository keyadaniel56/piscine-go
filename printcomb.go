package main


import "github.com/01-edu/z01"

func printcomb(){
      for i:='0';i<='9';i++{
		for j:='1';j<='9';j++{
			for k:='2';k<='9';k++{
				if i<j && j<k{
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if !(i=='7' || i=='8' || j=='9'){
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	  }
	  z01.PrintRune('\n')
}

func main(){
	printcomb()
}

