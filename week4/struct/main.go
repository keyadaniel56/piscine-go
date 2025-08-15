package main

import "fmt"

type Capitalizer struct{
	Text string
}

func isSeparator(r rune)bool{
	return r==' ' || r=='.' || r=='!' || r==',' || r=='?'
}

func isnumeric(n rune)bool{
	return n>=0 || n <='9'
}
func (c *Capitalizer) isCapitalizedProperly()bool{
	if len(c.Text)==0{
		return false
	}
	runes:=[]rune(c.Text)
	for i:=0;i<len(runes);i++{
		if i==0 || isSeparator(runes[i-1]){
			if runes[i]>='a' && runes[i]<='z'{
				return false
			}
		}else if i==0 || isSeparator(runes[i-1]) &&isnumeric(runes[i-1]){
			return true
		}
	}
	return true
}


func main(){
	c1:=Capitalizer{Text: "Hello! How are you?"}
	c2:=Capitalizer{Text: "Hello How Are You"}
	c3:=Capitalizer{Text: "Whats 4this 100K?"}
	c4:=Capitalizer{Text: "Whatsthis4"}
	c5:=Capitalizer{Text: "!!!!Whatsthis4"}
	c6:=Capitalizer{Text: ""}
	fmt.Println(c1.isCapitalizedProperly())
	fmt.Println(c2.isCapitalizedProperly())
	fmt.Println(c3.isCapitalizedProperly())
	fmt.Println(c4.isCapitalizedProperly())
	fmt.Println(c5.isCapitalizedProperly())
	fmt.Println(c6.isCapitalizedProperly())
}