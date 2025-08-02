package main

import "os"

func main() {
	if len(os.Args) < 1 {
		os.Stdout.Write([]byte("\n"))
	}
	last := os.Args[len(os.Args)-1]
	os.Stdout.Write([]byte(last + "\n"))
}
