package main

import "os"

func main(){
	last_args:=os.Args[len(os.Args)-1]
	os.Stdout.Write([]byte(last_args+"\n"))
	
}